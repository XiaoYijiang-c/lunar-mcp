package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// DestinyAnalysisTool returns destiny analysis based on bazi
var DestinyAnalysisTool = &Tool{
	Name:        "destiny_analysis",
	Description: "命理分析，基于八字分析五行强弱",
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

		// Get bazi and wuxing
		bazi := lunar.GetBaZi()
		wuxing := lunar.GetBaZiWuXing()

		// Count wuxing
		wuxingCount := map[string]int{
			"木": 0, "火": 0, "土": 0, "金": 0, "水": 0,
		}
		for _, w := range wuxing {
			wuxingCount[w]++
		}

		// Simple analysis based on day gan
		dayGan := lunar.GetDayGan()

		// Basic yongshen analysis (simplified)
		var yongshen string
		var analysis string

		switch dayGan {
		case "甲", "乙":
			yongshen = "火、土"
			analysis = "木性生发，喜火泄秀、土培根"
		case "丙", "丁":
			yongshen = "水、金"
			analysis = "火性炎上，喜水制、火金相生"
		case "戊", "己":
			yongshen = "木、金"
			analysis = "土性厚重，喜木疏土、金泄秀"
		case "庚", "辛":
			yongshen = "土、火"
			analysis = "金性肃杀，喜土生、火锻炼"
		case "壬", "癸":
			yongshen = "木、火"
			analysis = "水性润下，喜木泄秀、火温暖"
		default:
			yongshen = "待分析"
			analysis = "请咨询专业命理师"
		}

		return map[string]interface{}{
			"date":      fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"lunar":     lunar.ToFullString(),
			"dayGan":    dayGan,
			"dayZhi":    lunar.GetDayZhi(),
			"animal":    lunar.GetAnimal(),
			"bazi": map[string]string{
				"year":  bazi[0],
				"month": bazi[1],
				"day":   bazi[2],
				"time":  bazi[3],
			},
			"wuxing": map[string]int{
				"木": wuxingCount["木"],
				"火": wuxingCount["火"],
				"土": wuxingCount["土"],
				"金": wuxingCount["金"],
				"水": wuxingCount["水"],
			},
			"analysis": map[string]string{
				"yongshen": yongshen,
				"summary":  analysis,
			},
		}, nil
	},
}
