package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// EightCharFullTool returns complete eight character analysis
var EightCharFullTool = &Tool{
	Name:        "eight_char_full",
	Description: "获取完整的八字分析，包含十神、纳音五行",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "公历年份"},
			"month": map[string]interface{}{"type": "integer", "description": "公历月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "公历日期"},
			"hour":  map[string]interface{}{"type": "integer", "description": "时辰 (0-23，可选)"},
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

		// Get basic bazi
		bazi := lunar.GetBaZi()
		baziNaYin := lunar.GetBaZiNaYin()
		baziWuXing := lunar.GetBaZiWuXing()
		baziShiShenGan := lunar.GetBaZiShiShenGan()
		baziShiShenZhi := lunar.GetBaZiShiShenZhi()

		return map[string]interface{}{
			"date":      fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"lunar":     lunar.ToFullString(),
			"animal":    lunar.GetAnimal(),
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
			"wuxing": map[string]string{
				"year":  baziWuXing[0],
				"month": baziWuXing[1],
				"day":   baziWuXing[2],
				"time":  baziWuXing[3],
			},
			"shishen": map[string]string{
				"gan": baziShiShenGan[0],
				"zhi": baziShiShenZhi[0],
			},
			"dayGan":   lunar.GetDayGan(),
			"dayZhi":   lunar.GetDayZhi(),
			"dayNaYin": lunar.GetDayNaYin(),
			"fullBazi": fmt.Sprintf("%s %s %s %s", bazi[0], bazi[1], bazi[2], bazi[3]),
		}, nil
	},
}
