package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/example/lunar-mcp/internal/protocol"
	"github.com/example/lunar-mcp/internal/tools"
)

func main() {
	// Create registry and register tools
	registry := tools.NewRegistry()
	registry.Register(tools.LunarDateTool)
	registry.Register(tools.SolarTermsTool)
	registry.Register(tools.FestivalsTool)
	registry.Register(tools.AuspiciousDateTool)
	registry.Register(tools.DailyOmenTool)
	registry.Register(tools.ZodiacBaziTool)
	registry.Register(tools.SolarCalendarTool)
	registry.Register(tools.MonthCalendarTool)
	registry.Register(tools.YearCalendarTool)
	registry.Register(tools.EightCharFullTool)
	registry.Register(tools.DailyFortuneTool)
	registry.Register(tools.TimeBaziTool)
	registry.Register(tools.DestinyAnalysisTool)
	registry.Register(tools.TaoHolidayTool)
	registry.Register(tools.BuddhistHolidayTool)
	registry.Register(tools.LunarCalendarTool)
	registry.Register(tools.DateCalculatorTool)
	registry.Register(tools.LunarToSolarTool)
	registry.Register(tools.AuspiciousTimeTool)
	registry.Register(tools.DateSelectorTool)
	registry.Register(tools.MarriageCompatTool)
	registry.Register(tools.NameGeneratorTool)

	// Create handler
	handler := protocol.NewHandler()

	// Register initialize method
	handler.RegisterMethod("initialize", func(params map[string]interface{}) (interface{}, error) {
		return map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools": struct{}{},
			},
			"serverInfo": map[string]string{
				"name":    "lunar-mcp",
				"version": "1.0.0",
			},
		}, nil
	})

	// Register tools/list method
	handler.RegisterMethod("tools/list", func(params map[string]interface{}) (interface{}, error) {
		return registry.GetListResult(), nil
	})

	// Register tools/call method
	handler.RegisterMethod("tools/call", func(params map[string]interface{}) (interface{}, error) {
		name, hasName := params["name"].(string)
		if !hasName {
			return nil, fmt.Errorf("missing tool name")
		}

		args, hasArgs := params["arguments"].(map[string]interface{})
		if !hasArgs {
			args = make(map[string]interface{})
		}

		// Get tool from registry
		tool, ok := registry.Get(name)
		if !ok {
			return nil, fmt.Errorf("tool not found: %s", name)
		}

		// Call tool handler
		return tool.Handler(args)
	})

	// Register ping method
	handler.RegisterMethod("ping", func(params map[string]interface{}) (interface{}, error) {
		return map[string]string{"status": "pong"}, nil
	})

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// RPC endpoint
	http.HandleFunc("/rpc", handler.Handle)

	// Get port from env or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting lunar-mcp server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
