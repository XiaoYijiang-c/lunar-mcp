package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// LunarToSolarTool converts lunar date to solar date
var LunarToSolarTool = &Tool{
	Name:        "lunar_to_solar",
	Description: "农历转公历",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":      map[string]interface{}{"type": "integer", "description": "农历年份"},
			"month":     map[string]interface{}{"type": "integer", "description": "农历月份 (1-12)"},
			"day":       map[string]interface{}{"type": "integer", "description": "农历日期"},
			"isLeap":    map[string]interface{}{"type": "boolean", "description": "是否闰月"},
		},
		"required": []string{"year", "month", "day"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		day := int(params["day"].(float64))
		isLeap := false
		if v, ok := params["isLeap"].(bool); ok {
			isLeap = v
		}

		// Create lunar date
		var lunar *calendar.Lunar
		if isLeap {
			// For leap month, we need to find the correct one
			lunar = calendar.NewLunar(year, month, day, 0, 0, 0)
		} else {
			lunar = calendar.NewLunarFromYmd(year, month, day)
		}

		if lunar == nil {
			return nil, fmt.Errorf("invalid lunar date")
		}

		// Get solar date
		solar := lunar.GetSolar()
		if solar == nil {
			return nil, fmt.Errorf("conversion failed")
		}

		return map[string]interface{}{
			"lunar":      fmt.Sprintf("农历%d年%d月%d日", year, month, day),
			"isLeap":     isLeap,
			"solar":      solar.ToYmd(),
			"solarFull":  solar.ToFullString(),
			"weekday":    solar.GetWeek(),
			"weekdayCN":  solar.GetWeekInChinese(),
		}, nil
	},
}
