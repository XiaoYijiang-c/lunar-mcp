package protocol

// Error codes based on JSON-RPC 2.0 spec
const (
	ErrParseError      = -32700
	ErrInvalidRequest  = -32600
	ErrMethodNotFound  = -32601
	ErrInvalidParams   = -32602
	ErrInternalError   = -32603
)

// Error represents a JSON-RPC 2.0 error
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewError creates a new error
func NewError(code int, message string, data interface{}) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// Predefined errors
var (
	ErrParse        = NewError(ErrParseError, "Parse error", nil)
	ErrInvalidReq   = NewError(ErrInvalidRequest, "Invalid Request", nil)
	ErrNotFound     = NewError(ErrMethodNotFound, "Method not found", nil)
	ErrBadParams    = NewError(ErrInvalidParams, "Invalid params", nil)
	ErrInternal     = NewError(ErrInternalError, "Internal error", nil)
)
