package protocol

import "encoding/json"

// Response represents a JSON-RPC 2.0 response
type Response struct {
	JsonRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *Error     `json:"error,omitempty"`
	ID      interface{} `json:"id,omitempty"`
}

// NewSuccessResponse creates a successful response
func NewSuccessResponse(id interface{}, result interface{}) Response {
	return Response{
		JsonRPC: "2.0",
		Result:  result,
		ID:      id,
	}
}

// NewErrorResponse creates an error response
func NewErrorResponse(id interface{}, err *Error) Response {
	return Response{
		JsonRPC: "2.0",
		Error:   err,
		ID:      id,
	}
}

// MarshalJSON marshals the response to JSON
func (r Response) MarshalJSON() ([]byte, error) {
	// Custom marshaling to handle nil error
	if r.Error == nil {
		return json.Marshal(struct {
			JsonRPC string      `json:"jsonrpc"`
			Result  interface{} `json:"result,omitempty"`
			ID      interface{} `json:"id,omitempty"`
		}{
			JsonRPC: r.JsonRPC,
			Result:  r.Result,
			ID:      r.ID,
		})
	}
	return json.Marshal(r)
}
