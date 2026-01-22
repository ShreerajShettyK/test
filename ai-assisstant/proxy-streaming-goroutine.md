# Channel and Goroutine Usage for Streaming Responses:
Server-Sent Events (SSE) style streaming to send LLM responses to the client in real-time.
Flusher(from net/http package) is required to send data immediately without buffering

The goroutine and channel pattern in [services/apim/apim.go](services/apim/apim.go) (lines 64-77) provides **non-blocking, concurrent streaming** of responses. Here's how it improves performance:

## Key Performance Benefits

### 1. **Non-Blocking Return**
The function returns immediately with a channel, allowing the caller to start processing responses while data is still being received:

```go
// Returns immediately - doesn't wait for full response
responseStream := make(chan models.Response, 100)
return responseStream, nil  // Caller can start reading right away
```

### 2. **Concurrent Processing**
The goroutine processes the HTTP response body in the background while your application continues executing:

```go
go func() {
    // This runs concurrently - doesn't block the caller
    azure.ProcessStreamingResponse(resp.Body, responseStream)
}()
```

### 3. **Memory Efficiency**
The buffered channel (`make(chan models.Response, 100)`) acts as a queue that:
- Allows the producer (goroutine) to write up to 100 responses without blocking
- Prevents overwhelming the consumer with data
- Reduces memory usage by processing chunks incrementally instead of loading the entire response

### 4. **Resource Cleanup**
The `defer` statements ensure resources are properly released even if errors occur:

```go
defer close(responseStream)  // Signals completion to consumers
defer resp.Body.Close()      // Releases HTTP connection
```

## How It Works in Practice

Looking at how this is consumed in [services/common/handlers/response_handler.go](services/common/handlers/response_handler.go):

```go
func StreamResponseToClient(ctx context.Context, w http.ResponseWriter, responseStream <-chan models.Response) {
    for response := range responseStream {
        // Process each response as it arrives
        responseJSON, _ := json.Marshal(response)
        w.Write(append(responseJSON, '\n', '\n'))
        flusher.Flush()  // Send immediately to client
    }
}
```

## Performance Comparison

**Without Goroutine/Channel (Blocking):**
```
Client Request → Wait for full response → Process → Send to client
Total Time: Network + Processing + Send
```

**With Goroutine/Channel (Concurrent):**
```
Client Request → Start receiving
                 ↓
                 Receive chunk → Process → Send (concurrent)
                 ↓
                 Receive chunk → Process → Send (concurrent)
Total Time: Overlapped operations = Lower latency
```

## Real-World Impact

Based on the code structure across similar implementations ([services/azure/azure_stream_request.go](services/azure/azure_stream_request.go), [services/gemini/gemini_stream_request.go](services/gemini/gemini_stream_request.go), [services/aws/bedrock.go](services/aws/bedrock.go)), this pattern enables:

1. **Lower Time-To-First-Byte (TTFB)**: Client sees first response chunk faster
2. **Better Resource Utilization**: CPU processes while waiting for network I/O
3. **Scalability**: Server can handle multiple concurrent requests without blocking
4. **Graceful Error Handling**: Errors in processing don't leave connections hanging

The buffered channel size of 100 is a reasonable default that balances memory usage with throughput for typical LLM streaming responses.




Supporting code:

func StreamAzureAPIMResponse(ctx context.Context, wrappedRequest interface{}, stream bool) (<-chan models.Response, error) {
	// _, span, logger := lcpotel.Wrapper.StartTracer(ctx, lcpotel.Wrapper.GetCurrentPackageName(), lcpotel.Wrapper.GetCurrentMethodName())
	// defer span.End()

	logger.Infof("Starting StreamAzureAPIMResponse with streaming mode: %v", stream)

	cfg := utils.ApplicationProperties.AzureApim
	err := validateConfig(cfg)
	if err != nil {
		logger.Errorf("Failed to Validate Config")
		return nil, err
	}

	azureAPIMURL := fmt.Sprintf("%s/deployments/%s/chat/completions?api-version=%s",
		cfg.Apim_endpoint, cfg.Apim_deployment_name, cfg.Apim_deployment_version)
	logger.Infof("Azure APIM URL: %s", azureAPIMURL)

	req, err := prepareAPIMRequest(azureAPIMURL, wrappedRequest, cfg.Apim_key)
	if err != nil {
		logger.Errorf("Failed to prepare Azure APIM request: %v", err)
		return nil, err
	}
	logger.Infof("Azure APIM request prepared successfully")

	client := &http.Client{}
	logger.Infof("Sending request to Azure APIM")
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to send Azure APIM request: %v", err)
		return nil, fmt.Errorf("error making Azure APIM request: %w", err)
	}
	logger.Infof("Azure APIM request sent successfully, received response")

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		logger.Errorf("Azure APIM request failed with status code %d: %s", resp.StatusCode, respBody)
		return nil, fmt.Errorf("azure APIM request failed with status code %d: %s", resp.StatusCode, respBody)
	}
	logger.Infof("Azure APIM request succeeded with status code %d", resp.StatusCode)

	responseStream := make(chan models.Response, 100)
	logger.Infof("Starting response stream handler")

	go func() {
		defer close(responseStream)
		defer resp.Body.Close()

		if stream {
			logger.Infof("Processing streaming response from Azure APIM")
			azure.ProcessStreamingResponse(resp.Body, responseStream)
			logger.Infof("Completed processing streaming response")
		} else {
			logger.Infof("Processing non-streaming response from Azure APIM")
			azure.ProcessNonStreamingResponse(resp.Body, responseStream)
			logger.Infof("Completed processing non-streaming response")
		}
	}()

	return responseStream, nil
}


func handleCircuitBreaker(ctx context.Context, w http.ResponseWriter, req models.LLMRequest) error {
	logger.Infof("Starting handleCircuitBreaker for request: %+v", req)

	cbDAO := dao.NewCircuitBreakerDAO(config.DB)
	auditLogDAO := dao.NewAuditLogDAO(config.DB)
	cbManager := circuit.NewCircuitBreakerManager(cbDAO, auditLogDAO)

	responseStream, err := ProcessProvidersFn(ctx, req, cbManager)
	if err != nil {
		logger.Errorf("Failed to process providers for request: %+v, error: %v", req, err)
		SendErrorResponse(w, err, http.StatusInternalServerError)
		return err
	}

	logger.Infof("Streaming response to client for request: %+v", req)

	StreamResponseToClient(ctx, w, responseStream)
	logger.Infof("Completed streaming response for request: %+v", req)

	return nil
}

func StreamResponseToClient(ctx context.Context, w http.ResponseWriter, responseStream <-chan models.Response) {
	// _, span, logger := lcpotel.Wrapper.StartTracer(ctx, lcpotel.Wrapper.GetCurrentPackageName(), lcpotel.Wrapper.GetCurrentMethodName())
	// defer span.End()

	logger.Infof("Starting to stream response to client")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)
	flusher, ok := w.(http.Flusher)
	if !ok {
		logger.Errorf("ResponseWriter does not support streaming")
		return
	}
	for response := range responseStream {
		logger.Infof("Streaming response: %+v", response)
		responseJSON, err := json.Marshal(response)
		if err != nil {
			logger.Errorf("Error marshalling response to JSON: %v", err)
			continue
		}
		_, err = w.Write(append(responseJSON, '\n', '\n'))
		if err != nil {
			logger.Errorf("Error writing response to client: %v", err)
			return
		}
		flusher.Flush()
	}

	logger.Infof("Finished streaming all responses to client")
}


Key Concept: http.Flusher
The flusher.Flush() call is critical:

Without Flush:

Response chunks → Server buffer → Wait for buffer to fill → Send to client
(Client experiences delays)

With Flush:

Response chunks → Immediately sent to client
(Client sees data as soon as it's available)

This creates the typewriter effect users see when chatting with LLMs - text appears progressively rather than all at once.

