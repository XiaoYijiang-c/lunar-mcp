package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// AdvancedBaziTool returns detailed bazi analysis with shishen
var AdvancedBaziTool = &Tool{
	Name:        "advanced_bazi",
	Description: "进阶八字分析，包含完整十神、大运流年",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "年份"},
			"month": map[string]interface{}{"type": "integer", "description": "月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "日期"},
			"hour":  map[string]interface{}{"type": "integer", "description": "时辰(0-23)"},
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

		// Get shishen (十神)
		shishenGan := lunar.GetBaZiShiShenGan()
		shishenZhi := lunar.GetBaZiShiShenZhi()

		// Get day gan/zhi for self
		dayGan := lunar.GetDayGan()
		dayZhi := lunar.GetDayZhi()

		// Get detailed info for each position
		positions := []string{"year", "month", "day", "time"}

		var baziDetail []map[string]interface{}
		for i, pos := range positions {
			baziDetail = append(baziDetail, map[string]interface{}{
				"position":   pos,
				"ganZhi":    bazi[i],
				"wuxing":    baziWuXing[i],
				"naYin":     baziNaYin[i],
				"shishenGan": shishenGan[i],
				"shishenZhi": shishenZhi[i],
			})
		}

		// Get year/month/day zodiac
		yearZhi := lunar.GetYearZhi()
		monthZhi := lunar.GetMonthZhi()

		// Calculate dachen (年柱空亡)
		var dachen string
		xunKong := lunar.GetYearXunKong()
		if xunKong != "" {
			dachen = xunKong + "空"
		}

		return map[string]interface{}{
			"date":      solar.ToYmd(),
			"lunar":     lunar.ToFullString(),
			"animal":    lunar.GetAnimal(),
			"dayGan":    dayGan,
			"dayZhi":    dayZhi,
			"dayNaYin":  lunar.GetDayNaYin(),
			"dayWuxing": baziWuXing[2],
			"bazi":      baziDetail,
			"dachen":    dachen,
			"yearZhi":   yearZhi,
			"monthZhi":  monthZhi,
		}, nil
	},
}
