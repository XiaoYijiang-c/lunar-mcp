package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// DateCalculatorTool calculates dates (forward/backward)
var DateCalculatorTool = &Tool{
	Name:        "date_calculator",
	Description: "日期推算，支持向前或向后计算日期",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "起始年份"},
			"month": map[string]interface{}{"type": "integer", "description": "起始月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "起始日期"},
			"days":  map[string]interface{}{"type": "integer", "description": "偏移天数，正数向后，负数向前"},
		},
		"required": []string{"year", "month", "day", "days"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		day := int(params["day"].(float64))
		days := int(params["days"].(float64))

		solar := calendar.NewSolarFromYmd(year, month, day)
		if solar == nil {
			return nil, fmt.Errorf("invalid date")
		}

		// Calculate new date
		nextSolar := solar.NextDay(days)

		lunar := calendar.NewLunarFromSolar(nextSolar)
		if lunar == nil {
			return nil, fmt.Errorf("invalid lunar date")
		}

		return map[string]interface{}{
			"startDate":  fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"daysOffset": days,
			"resultDate": nextSolar.ToYmd(),
			"lunar":     lunar.ToFullString(),
			"weekday":   nextSolar.GetWeek(),
		}, nil
	},
}
