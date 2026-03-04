package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// SolarTermsTool returns solar terms information
var SolarTermsTool = &Tool{
	Name:        "solar_terms",
	Description: "获取指定日期的节气信息",
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

		// Get current jie qi
		currentJie := lunar.GetCurrentJie()
		currentJieQi := lunar.GetCurrentJieQi()

		return map[string]interface{}{
			"date":           fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"currentJie":    currentJie.GetName(),
			"currentJieQi":  currentJieQi.GetName(),
			"jieQiList":     lunar.GetJieQiList(),
			"jieQiTable":    lunar.GetJieQiTable(),
		}, nil
	},
}
