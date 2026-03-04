package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// BuddhistHolidayTool returns Buddhist holidays
var BuddhistHolidayTool = &Tool{
	Name:        "buddhist_holiday",
	Description: "获取佛历相关信息",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "公历年份"},
			"month": map[string]interface{}{"type": "integer", "description": "公历月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "公历日期"},
		},
		"required": []string{"year", "month", "day"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		day := int(params["day"].(float64))

		solar := calendar.NewSolarFromYmd(year, month, day)
		if solar == nil {
			return nil, fmt.Errorf("invalid date")
		}

		lunar := calendar.NewLunarFromSolar(solar)
		if lunar == nil {
			return nil, fmt.Errorf("invalid lunar date")
		}

		// Get Foto (佛历)
		foto := lunar.GetFoto()

		// Calculate Buddhist year (佛历) - starts from 1027 BC
		buddhistYear := year + 1027

		return map[string]interface{}{
			"date":            fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"buddhistYear":    buddhistYear,
			"lunar":           lunar.ToFullString(),
			"buddhistInfo": map[string]interface{}{
				"foto": foto,
			},
		}, nil
	},
}
