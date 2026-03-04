package protocol

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Handler handles JSON-RPC 2.0 requests
type Handler struct {
	methods map[string]MethodHandler
}

// MethodHandler is a function that handles a method
type MethodHandler func(params map[string]interface{}) (interface{}, error)

// NewHandler creates a new handler
func NewHandler() *Handler {
	return &Handler{
		methods: make(map[string]MethodHandler),
	}
}

// RegisterMethod registers a method handler
func (h *Handler) RegisterMethod(name string, fn MethodHandler) {
	h.methods[name] = fn
}

// Handle handles an HTTP request
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Only allow POST
	if r.Method != http.MethodPost {
		resp := NewErrorResponse(nil, ErrInvalidReq)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp := NewErrorResponse(nil, ErrParse)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Parse request
	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("Parse error: %v", err)
		resp := NewErrorResponse(nil, ErrParse)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Handle batch requests (simplified - just handle single for now)
	h.handleRequest(w, req)
}

func (h *Handler) handleRequest(w http.ResponseWriter, req Request) {
	// Get method handler
	handler, ok := h.methods[req.Method]
	if !ok {
		resp := NewErrorResponse(req.ID, ErrNotFound)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Parse params
	var params map[string]interface{}
	if len(req.Params) > 0 {
		if err := json.Unmarshal(req.Params, &params); err != nil {
			resp := NewErrorResponse(req.ID, ErrBadParams)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}

	// Call handler
	result, err := handler(params)
	if err != nil {
		resp := NewErrorResponse(req.ID, NewError(ErrInternalError, err.Error(), nil))
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Success
	resp := NewSuccessResponse(req.ID, result)
	json.NewEncoder(w).Encode(resp)
}
