package protocol

import "time"

// Error codes based on JSON-RPC 2.0 spec
const (
	ErrParseError     = -32700 // Parse error
	ErrInvalidRequest = -32600 // Invalid Request
	ErrMethodNotFound = -32601 // Method not found
	ErrInvalidParams  = -32602 // Invalid params
	ErrInternalError = -32603 // Internal error

	// Custom error codes (>= -32000)
	ErrToolNotFound     = -32001 // Tool not found
	ErrToolExecution   = -32002 // Tool execution failed
	ErrSessionExpired  = -32003 // Session expired
	ErrAuthFailed     = -32004 // Authentication failed
	ErrRateLimited    = -32005 // Rate limited
	ErrTimeout        = -32006 // Request timeout
)

// Error represents a JSON-RPC 2.0 error
type Error struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time  `json:"timestamp"`
	RequestID string     `json:"requestId,omitempty"`
}

// NewError creates a new error with timestamp
func NewError(code int, message string, data interface{}) *Error {
	return &Error{
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewErrorWithRequest creates error with request ID for tracing
func NewErrorWithRequest(code int, message string, requestID string, data interface{}) *Error {
	return &Error{
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
		RequestID: requestID,
	}
}

// Predefined errors
var (
	ErrParse       = NewError(ErrParseError, "Parse error", nil)
	ErrInvalidReq  = NewError(ErrInvalidRequest, "Invalid Request", nil)
	ErrNotFound    = NewError(ErrMethodNotFound, "Method not found", nil)
	ErrBadParams   = NewError(ErrInvalidParams, "Invalid params", nil)
	ErrInternal    = NewError(ErrInternalError, "Internal error", nil)

	ErrToolNotFoundErr  = NewError(ErrToolNotFound, "Tool not found", nil)
	ErrToolExec      = NewError(ErrToolExecution, "Tool execution failed", nil)
	ErrSessionExp    = NewError(ErrSessionExpired, "Session expired", nil)
	ErrAuth         = NewError(ErrAuthFailed, "Authentication failed", nil)
	ErrRateLimit    = NewError(ErrRateLimited, "Rate limited", nil)
	ErrTimeoutError = NewError(ErrTimeout, "Request timeout", nil)
)

// ErrorCodeToMessage returns human readable message for error code
func ErrorCodeToMessage(code int) string {
	messages := map[int]string{
		ErrParseError:     "Parse error - Invalid JSON",
		ErrInvalidRequest: "Invalid Request - Missing or invalid method",
		ErrMethodNotFound: "Method not found",
		ErrInvalidParams:  "Invalid params - Check your input",
		ErrInternalError:  "Internal error - Server issue",

		ErrToolNotFound:    "The requested tool does not exist",
		ErrToolExecution:   "Tool execution failed",
		ErrSessionExpired:  "Session has expired, please reinitialize",
		ErrAuthFailed:      "Authentication failed",
		ErrRateLimited:     "Too many requests, please slow down",
		ErrTimeout:         "Request timeout",
	}
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "Unknown error"
}
