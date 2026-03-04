package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// TaoHolidayTool returns Taoist holidays
var TaoHolidayTool = &Tool{
	Name:        "tao_holiday",
	Description: "获取道历相关信息",
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

		// Get Tao info
		tao := lunar.GetTao()

		// Calculate Dao Lü (道历) - starts from 2697 BC
		daoLiYear := year + 2697

		return map[string]interface{}{
			"date":        fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"daoLiYear":   daoLiYear,
			"lunar":       lunar.ToFullString(),
			"taoInfo": map[string]interface{}{
				"tao":    tao,
			},
		}, nil
	},
}
