package tests

import (
	"testing"

	"github.com/example/lunar-mcp/internal/tools"
)

// Test: Registry tools exist
func TestRegistryTools(t *testing.T) {
	registry := tools.NewRegistry()
	registry.Register(tools.LunarDateTool)
	registry.Register(tools.ZodiacBaziTool)
	registry.Register(tools.SolarTermsTool)
	registry.Register(tools.FestivalsTool)
	registry.Register(tools.AuspiciousDateTool)
	registry.Register(tools.DailyOmenTool)
	registry.Register(tools.SolarCalendarTool)
	registry.Register(tools.MonthCalendarTool)
	registry.Register(tools.YearCalendarTool)
	registry.Register(tools.EightCharFullTool)
	registry.Register(tools.DestinyAnalysisTool)
	registry.Register(tools.DailyFortuneTool)
	registry.Register(tools.TimeBaziTool)
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

	list := registry.List()
	if len(list) < 28 {
		t.Errorf("Expected at least 28 tools, got %d", len(list))
	}

	t.Logf("Total tools registered: %d", len(list))
}

// Test: Tool names
func TestToolNames(t *testing.T) {
	expected := []string{
		"lunar_date",
		"zodiac_bazi",
		"solar_terms",
		"festivals",
		"auspicious_date",
		"daily_omen",
	}

	registry := tools.NewRegistry()
	for _, name := range expected {
		registry.Register(&tools.Tool{Name: name})
	}

	list := registry.List()
	if len(list) != len(expected) {
		t.Errorf("Expected %d tools, got %d", len(expected), len(list))
	}
}
