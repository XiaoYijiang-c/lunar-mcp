package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/example/lunar-mcp/internal/protocol"
	"github.com/example/lunar-mcp/internal/session"
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
	registry.Register(tools.IChingDivinationTool)
	registry.Register(tools.NineStarFlyingTool)
	registry.Register(tools.AdvancedBaziTool)
	registry.Register(tools.PengzuBaijiTool)
	registry.Register(tools.FortunePeriodsTool)
	registry.Register(tools.ShenShaTool)

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

	// Register tools/register method (dynamic tool)
	handler.RegisterMethod("tools/register", func(params map[string]interface{}) (interface{}, error) {
		name, _ := params["name"].(string)
		description, _ := params["description"].(string)
		inputSchema, _ := params["inputSchema"].(map[string]interface{})
		
		err := registry.RegisterDynamic(tools.DynamicToolRequest{
			Name:        name,
			Description: description,
			InputSchema: inputSchema,
			Handler: func(p map[string]interface{}) (interface{}, error) {
				return map[string]string{"status": "dynamic tool called"}, nil
			},
		})
		
		if err != nil {
			return nil, err
		}
		return map[string]string{"status": "registered", "tool": name}, nil
	})

	// Register tools/unregister method
	handler.RegisterMethod("tools/unregister", func(params map[string]interface{}) (interface{}, error) {
		name, _ := params["name"].(string)
		removed := registry.Unregister(name)
		return map[string]interface{}{"status": removed, "tool": name}, nil
	})

	// Register session management
	sessionMgr := session.NewManager()

	// Health check endpoint with metrics
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "ok",
			"uptime":   time.Since(startTime).Seconds(),
			"sessions": sessionMgr.Count(),
		})
	})

	// Metrics endpoint
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ServerMetrics.GetMetrics())
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
