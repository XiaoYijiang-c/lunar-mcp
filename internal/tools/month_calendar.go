package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// MonthCalendarTool returns a full month calendar
var MonthCalendarTool = &Tool{
	Name:        "month_calendar",
	Description: "获取指定月份的完整日历",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "公历年份"},
			"month": map[string]interface{}{"type": "integer", "description": "公历月份"},
		},
		"required": []string{"year", "month"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))

		// Get days in month
		var days []map[string]interface{}
		
		for day := 1; day <= 31; day++ {
			solar := calendar.NewSolarFromYmd(year, month, day)
			if solar == nil {
				break
			}

			lunar := calendar.NewLunarFromSolar(solar)
			if lunar == nil {
				continue
			}

			// Get solar terms
			var jieQi string
			if currentJie := lunar.GetCurrentJie(); currentJie != nil {
				jieQi = currentJie.GetName()
			}

			days = append(days, map[string]interface{}{
				"date":       fmt.Sprintf("%d-%02d-%02d", year, month, day),
				"weekday":    solar.GetWeek(),
				"weekdayCN":  solar.GetWeekInChinese(),
				"lunar":      lunar.ToFullString(),
				"animal":     lunar.GetAnimal(),
				"solarTerms": jieQi,
			})
		}

		return map[string]interface{}{
			"year":  year,
			"month": month,
			"days":  days,
			"total": len(days),
		}, nil
	},
}
