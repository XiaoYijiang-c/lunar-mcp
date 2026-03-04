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
