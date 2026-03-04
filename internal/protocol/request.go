package protocol

import "encoding/json"

// Request represents a JSON-RPC 2.0 request
type Request struct {
	JsonRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      interface{}     `json:"id,omitempty"`
}

// ParseParams parses the params field into the given struct
func (r *Request) ParseParams(v interface{}) error {
	if len(r.Params) == 0 {
		return nil
	}
	return json.Unmarshal(r.Params, v)
}
