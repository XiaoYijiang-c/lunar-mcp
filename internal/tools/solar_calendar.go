package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// SolarCalendarTool returns detailed solar calendar information
var SolarCalendarTool = &Tool{
	Name:        "solar_calendar",
	Description: "获取公历详细信息，包括星座、星期、儒略日等",
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

		return map[string]interface{}{
			"date":          fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"weekday":       solar.GetWeek(),
			"weekdayCN":     solar.GetWeekInChinese(),
			"constellation": solar.GetXingZuo(),
			"julianDay":    solar.GetJulianDay(),
			"isLeapYear":   solar.IsLeapYear(),
			"fullString":    solar.ToFullString(),
		}, nil
	},
}
