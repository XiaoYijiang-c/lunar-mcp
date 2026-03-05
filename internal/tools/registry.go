package tools

import "encoding/json"

// Tool represents an MCP tool
type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"inputSchema"`
	Handler     func(params map[string]interface{}) (interface{}, error)
}

// Registry is the tool registry
type Registry struct {
	tools map[string]*Tool
}

// NewRegistry creates a new registry
func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]*Tool),
	}
}

// Register registers a tool
func (r *Registry) Register(tool *Tool) {
	r.tools[tool.Name] = tool
}

// Get returns a tool by name
func (r *Registry) Get(name string) (*Tool, bool) {
	tool, ok := r.tools[name]
	return tool, ok
}

// List returns all tools
func (r *Registry) List() []*Tool {
	result := make([]*Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		result = append(result, tool)
	}
	return result
}

// ListToolsResult represents the result of tools/list
type ListToolsResult struct {
	Tools []ToolInfo `json:"tools"`
}

// ToolInfo represents tool info for the list response
type ToolInfo struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema json.RawMessage        `json:"inputSchema"`
}

// GetListResult returns the result for tools/list
func (r *Registry) GetListResult() ListToolsResult {
	tools := make([]ToolInfo, 0, len(r.tools))
	for _, tool := range r.tools {
		schemaBytes, _ := json.Marshal(tool.InputSchema)
		tools = append(tools, ToolInfo{
			Name:        tool.Name,
			Description: tool.Description,
			InputSchema: schemaBytes,
		})
	}
	return ListToolsResult{Tools: tools}
}

// DynamicToolRequest represents a request to register a new tool
type DynamicToolRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"inputSchema"`
	Handler     func(params map[string]interface{}) (interface{}, error)
}

// RegisterDynamic registers a tool at runtime
func (r *Registry) RegisterDynamic(req DynamicToolRequest) error {
	if req.Name == "" {
		return ErrToolNameRequired
	}
	if req.Handler == nil {
		return ErrToolHandlerRequired
	}

	tool := &Tool{
		Name:        req.Name,
		Description: req.Description,
		InputSchema: req.InputSchema,
		Handler:     req.Handler,
	}

	r.tools[req.Name] = tool
	return nil
}

// Unregister removes a tool at runtime
func (r *Registry) Unregister(name string) bool {
	if _, ok := r.tools[name]; ok {
		delete(r.tools, name)
		return true
	}
	return false
}

// ErrToolNameRequired error
var ErrToolNameRequired = &ToolError{"tool name is required"}

// ErrToolHandlerRequired error
var ErrToolHandlerRequired = &ToolError{"tool handler is required"}

// ToolError represents a tool error
type ToolError struct {
	Message string
}

func (e *ToolError) Error() string {
	return e.Message
}
