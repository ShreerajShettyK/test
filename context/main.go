package main

import (
	"fmt"
	"net/http"
	"time"
)

// Context flows down the call chain
// Used for request metadata (user ID, request ID, auth info)
// Using context.WithValue Avoids global variables
// func main() {
// 	ctx := context.Background()
// 	newCtx := context.WithValue(ctx, "api-key", "sdyuiodsdfghjklkjhgf")

// 	processRequest(newCtx)
// }

// func processRequest(ctx context.Context) {
// 	apiKey := ctx.Value("api-key")
// 	fmt.Println(apiKey)
// }

// ✔ Use context values only for request-scoped data
// ✔ Do NOT store large data or optional parameters
// ✔ Contexts are immutable (each WithValue creates a new one)

// Context for cancellation and timeouts
// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel() // always call cancel to free resources

// 	select {
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Work completed")
// 	case <-ctx.Done():
// 		fmt.Println("Context cancelled:", ctx.Err())
// 	}
// }

// How HTTP Frameworks Use Context (Very Important)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // request-scoped context

	fmt.Println("Request started")

	select {
	case <-time.After(10 * time.Second):
		// Simulate long work
		fmt.Fprintln(w, "Request completed successfully")

	case <-ctx.Done():
		// Triggered when client disconnects
		fmt.Println("Client cancelled request:", ctx.Err())
		return
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on http://localhost:8084")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
