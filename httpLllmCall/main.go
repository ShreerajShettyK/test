package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type request struct {
	Question string `json:"question"`
}

const port = 8084

func main() {
	http.HandleFunc("/makeCall", handlerfunc)
	log.Println("Server started running on:", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Only Post allowed // content type not json", http.StatusMethodNotAllowed)
		return
	}

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Println("Received request at /makeCall")

	modelPayload := map[string]interface{}{
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": req.Question,
			},
		},
		"max_tokens":  4096,
		"temperature": 1,
		"top_p":       1,
	}

	payloadBytes, err := json.Marshal(modelPayload)
	if err != nil {
		http.Error(w, "Failed to marshal payload", http.StatusInternalServerError)
		return
	}

	reqAI, err := http.NewRequest("POST", "https://azure-open-ai-llm-client.openai.azure.com/openai/deployments/gpt-4o-regional-001/chat/completions?api-version=2023-03-15-preview",
		bytes.NewBuffer(payloadBytes))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	reqAI.Header.Set("Content-Type", "application/json")
	reqAI.Header.Set("Authorization", "Bearer efa8f572b4654f109db93aed83a11ba3")

	client := &http.Client{}
	resp, err := client.Do(reqAI)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
	fmt.Println("API call successfull")
}
