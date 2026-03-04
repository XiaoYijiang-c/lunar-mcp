package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// TimeBaziTool returns time-based bazi (八字 with specific hour)
var TimeBaziTool = &Tool{
	Name:        "time_bazi",
	Description: "获取指定时辰的八字，包含时辰对应的人生",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "公历年份"},
			"month": map[string]interface{}{"type": "integer", "description": "公历月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "公历日期"},
			"hour":  map[string]interface{}{"type": "integer", "description": "小时 (0-23)"},
			"minute": map[string]interface{}{"type": "integer", "description": "分钟 (0-59，可选)"},
		},
		"required": []string{"year", "month", "day", "hour"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		day := int(params["day"].(float64))
		hour := int(params["hour"].(float64))
		minute := 0
		if m, ok := params["minute"].(float64); ok {
			minute = int(m)
		}

		solar := calendar.NewSolar(year, month, day, hour, minute, 0)
		if solar == nil {
			return nil, fmt.Errorf("invalid date")
		}

		lunar := calendar.NewLunarFromSolar(solar)
		if lunar == nil {
			return nil, fmt.Errorf("invalid lunar date")
		}

		// Get bazi
		bazi := lunar.GetBaZi()
		baziNaYin := lunar.GetBaZiNaYin()

		// Get time specific info
		lunarTime := lunar.GetTime()
		timeGan := lunar.GetTimeGan()
		timeZhi := lunar.GetTimeZhi()
		timeNaYin := lunar.GetTimeNaYin()
		timeShengXiao := lunar.GetTimeShengXiao()

		_ = lunarTime // suppress unused warning

		return map[string]interface{}{
			"date":     fmt.Sprintf("%d-%02d-%02d %02d:%02d", year, month, day, hour, minute),
			"lunar":    lunar.ToFullString(),
			"bazi": map[string]string{
				"year":  bazi[0],
				"month": bazi[1],
				"day":   bazi[2],
				"time":  bazi[3],
			},
			"nayin": map[string]string{
				"year":  baziNaYin[0],
				"month": baziNaYin[1],
				"day":   baziNaYin[2],
				"time":  baziNaYin[3],
			},
			"timeInfo": map[string]string{
				"timeGan":    timeGan,
				"timeZhi":    timeZhi,
				"timeNaYin": timeNaYin,
				"animal":     timeShengXiao,
			},
			"fullBazi": fmt.Sprintf("%s %s %s %s", bazi[0], bazi[1], bazi[2], bazi[3]),
		}, nil
	},
}
