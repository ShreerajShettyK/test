package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LLMRequest struct {
	Prompt string `json:"prompt"`
}

type LLMResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method other than post not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Request triggered")

	var llmrequest LLMRequest
	err := json.NewDecoder(r.Body).Decode(&llmrequest)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	resp := LLMResponse{
		Message: "Received prompt: " + llmrequest.Prompt,
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/callLlm", handler)
	fmt.Println("Server running on http://localhost:8084")
	http.ListenAndServe(":8084", nil)
}
