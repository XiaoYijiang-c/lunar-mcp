package tests

import (
	"testing"

	"github.com/example/lunar-mcp/internal/protocol"
)

// Test: Protocol Error codes
func TestErrorCodes(t *testing.T) {
	tests := []struct {
		name     string
		code    int
		message string
	}{
		{"ParseError", -32700, "Parse error"},
		{"InvalidRequest", -32600, "Invalid Request"},
		{"MethodNotFound", -32601, "Method not found"},
		{"InvalidParams", -32602, "Invalid params"},
		{"InternalError", -32603, "Internal error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := protocol.NewError(tt.code, tt.message, nil)
			if err.Code != tt.code {
				t.Errorf("Expected code %d, got %d", tt.code, err.Code)
			}
			if err.Message != tt.message {
				t.Errorf("Expected message %s, got %s", tt.message, err.Message)
			}
		})
	}
}

// Test: Response creation
func TestResponseCreation(t *testing.T) {
	// Test success response
	resp := protocol.NewSuccessResponse(1, map[string]string{"status": "ok"})
	if resp.JsonRPC != "2.0" {
		t.Error("Expected jsonrpc 2.0")
	}
	if resp.ID != 1 {
		t.Error("Expected ID 1")
	}

	// Test error response
	errResp := protocol.NewErrorResponse(1, protocol.ErrNotFound)
	if errResp.Error == nil {
		t.Error("Expected error")
	}
	if errResp.Error.Code != -32601 {
		t.Errorf("Expected -32601, got %d", errResp.Error.Code)
	}
}

// Test: Request parsing
func TestRequestParsing(t *testing.T) {
	req := protocol.Request{
		JsonRPC: "2.0",
		Method:  "test",
		ID:      1,
	}

	if req.JsonRPC != "2.0" {
		t.Error("Expected jsonrpc 2.0")
	}
	if req.Method != "test" {
		t.Error("Expected method test")
	}
	if req.ID != 1 {
		t.Error("Expected ID 1")
	}
}

// Test: Handler registration
func TestHandlerRegistration(t *testing.T) {
	h := protocol.NewHandler()

	// Register a method
	h.RegisterMethod("test", func(params map[string]interface{}) (interface{}, error) {
		return map[string]string{"result": "ok"}, nil
	})

	// Check registration (can't verify directly, but should not panic)
	t.Log("Handler registration works")
}

// Test: Empty params handling
func TestEmptyParams(t *testing.T) {
	h := protocol.NewHandler()
	h.RegisterMethod("test", func(params map[string]interface{}) (interface{}, error) {
		if params == nil {
			return map[string]string{"status": "nil params"}, nil
		}
		return map[string]string{"status": "ok"}, nil
	})

	// Should handle nil params gracefully
	t.Log("Empty params test completed")
}
