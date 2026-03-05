package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SSEHandler handles Server-Sent Events
type SSEHandler struct {
	toolHandler http.Handler
}

// NewSSEHandler creates a new SSE handler
func NewSSEHandler(toolHandler http.Handler) *SSEHandler {
	return &SSEHandler{toolHandler: toolHandler}
}

// ServeHTTP handles SSE requests
func (h *SSEHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")

	// Notify client that connection is established
	fmt.Fprintf(w, "event: connected\ndata: {\"status\":\"connected\",\"time\":\"%s\"}\n\n", time.Now().Format(time.RFC3339))
	w.(http.Flusher).Flush()

	// Send message after 3 seconds as a demo
	go func() {
		time.Sleep(3 * time.Second)
		message := fmt.Sprintf("{\"message\":\"SSE connection active\",\"timestamp\":\"%s\"}", time.Now().Format(time.RFC3339))
		fmt.Fprintf(w, "event: message\ndata: %s\n\n", message)
		w.(http.Flusher).Flush()
	}()

	// Keep connection alive until closed
	<-r.Context().Done()
	fmt.Fprintf(w, "event: closed\ndata: {\"status\":\"disconnected\"}\n\n")
}

// Send sends an SSE message
func Send(w http.ResponseWriter, event string, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	if event != "" {
		fmt.Fprintf(w, "event: %s\n", event)
	}
	fmt.Fprintf(w, "data: %s\n\n", dataBytes)
	
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
	
	return nil
}

// SendProgress sends progress update
func SendProgress(w http.ResponseWriter, progress int, message string) {
	Send(w, "progress", map[string]interface{}{
		"progress": progress,
		"message":  message,
	})
}

// SendResult sends final result
func SendResult(w http.ResponseWriter, result interface{}) {
	Send(w, "result", result)
}

// SendError sends error
func SendError(w http.ResponseWriter, errMsg string) {
	Send(w, "error", map[string]string{"error": errMsg})
}
