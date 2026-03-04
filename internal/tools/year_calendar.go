package tools

import (
	"github.com/6tail/lunar-go/calendar"
)

// YearCalendarTool returns a full year calendar overview
var YearCalendarTool = &Tool{
	Name:        "year_calendar",
	Description: "获取指定年份的日历概览",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year": map[string]interface{}{"type": "integer", "description": "公历年份"},
		},
		"required": []string{"year"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))

		// Get year info for each month
		var months []map[string]interface{}

		for month := 1; month <= 12; month++ {
			// Get first day of month
			solar := calendar.NewSolarFromYmd(year, month, 1)
			if solar == nil {
				continue
			}

			lunar := calendar.NewLunarFromSolar(solar)
			if lunar == nil {
				continue
			}

			// Get last day of month
			lastDay := 31
			for d := 31; d >= 28; d-- {
				if calendar.NewSolarFromYmd(year, month, d) != nil {
					lastDay = d
					break
				}
			}

			months = append(months, map[string]interface{}{
				"month":         month,
				"firstDayLunar": lunar.ToFullString(),
				"lastDay":       lastDay,
			})
		}

		return map[string]interface{}{
			"year":   year,
			"months": months,
		}, nil
	},
}
