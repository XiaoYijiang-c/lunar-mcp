package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// LunarDateTool returns lunar calendar information for a given solar date
var LunarDateTool = &Tool{
	Name:        "lunar_date",
	Description: "获取指定公历日期的农历信息",
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

		return map[string]interface{}{
			"lunar":    lunar.ToFullString(),
			"solar":   solar.ToFullString(),
			"weekday": lunar.GetWeek(),
			"animal":  lunar.GetAnimal(),
		}, nil
	},
}
