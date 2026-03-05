package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// FortunePeriodsTool returns dachun and liunian (大运流年)
var FortunePeriodsTool = &Tool{
	Name:        "fortune_periods",
	Description: "大运流年分析",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "年份"},
			"month": map[string]interface{}{"type": "integer", "description": "月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "日期"},
			"gender": map[string]interface{}{"type": "string", "description": "性别: 男/女"},
		},
		"required": []string{"year", "month", "day", "gender"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		day := int(params["day"].(float64))
		gender := params["gender"].(string)

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
		dayGan := lunar.GetDayGan()
		dayZhi := lunar.GetDayZhi()

		// Determine yin/yang
		isYang := func(gan string) bool {
			yangGan := []string{"甲", "丙", "戊", "庚", "壬"}
			for _, g := range yangGan {
				if g == gan {
					return true
				}
			}
			return false
		}

		isMale := gender == "男"

		// Calculate starting dayun (simplified)
		// Get current year
		currentYear := 2026

		// Generate 10 big fortune periods (大运)
		var daYun []map[string]interface{}

		// Simplified dayun calculation based on month zhi
		monthZhi := lunar.GetMonthZhi()

		zhiOrder := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
		startIdx := 0
		for i, z := range zhiOrder {
			if z == monthZhi {
				startIdx = i
				break
			}
		}

		// Calculate next 10 dayun
		ganOrder := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

		dayGanIdx := 0
		for i, g := range ganOrder {
			if g == dayGan {
				dayGanIdx = i
				break
			}
		}

		// Determine direction based on gender and day gan yin/yang
		step := 1
		if isMale {
			if isYang(dayGan) {
				step = 1
			} else {
				step = -1
			}
		} else {
			if isYang(dayGan) {
				step = -1
			} else {
				step = 1
			}
		}

		ganIdx := dayGanIdx

		for i := 0; i < 10; i++ {
			zhiIdx := (startIdx + i + 1) % 12
			ganIdx = (ganIdx + step + 10) % 10

			age := 1 + i*10

			daYun = append(daYun, map[string]interface{}{
				"index":    i + 1,
				"age":      fmt.Sprintf("%d-%d岁", age, age+9),
				"ganZhi":   ganOrder[ganIdx] + zhiOrder[zhiIdx],
				"year":     currentYear + i*10,
			})
		}

		// Generate next 10 liunian (流年)
		var liuNian []map[string]interface{}
		for i := 0; i < 10; i++ {
			zhiIdx := (startIdx + i) % 12
			ganIdx = (dayGanIdx + i + 1) % 10

			liuNian = append(liuNian, map[string]interface{}{
				"year":    currentYear + i,
				"ganZhi": ganOrder[ganIdx] + zhiOrder[zhiIdx],
			})
		}

		return map[string]interface{}{
			"date":    solar.ToYmd(),
			"bazi":    bazi,
			"dayGan":  dayGan,
			"dayZhi":  dayZhi,
			"gender":  gender,
			"isYang":  isYang(dayGan),
			"daYun":   daYun,
			"liuNian": liuNian,
		}, nil
	},
}
