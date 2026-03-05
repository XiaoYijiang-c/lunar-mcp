package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// MarriageCompatTool checks marriage compatibility between two people
var MarriageCompatTool = &Tool{
	Name:        "marriage_compat",
	Description: "合婚，分析两人八字合婚情况",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"person1": map[string]interface{}{
				"type": "object",
				"description": "第一人",
				"properties": map[string]interface{}{
					"year":  map[string]interface{}{"type": "integer", "description": "年份"},
					"month": map[string]interface{}{"type": "integer", "description": "月份"},
					"day":   map[string]interface{}{"type": "integer", "description": "日期"},
					"name":  map[string]interface{}{"type": "string", "description": "姓名(可选)"},
				},
				"required": []string{"year", "month", "day"},
			},
			"person2": map[string]interface{}{
				"type": "object",
				"description": "第二人",
				"properties": map[string]interface{}{
					"year":  map[string]interface{}{"type": "integer", "description": "年份"},
					"month": map[string]interface{}{"type": "integer", "description": "月份"},
					"day":   map[string]interface{}{"type": "integer", "description": "日期"},
					"name":  map[string]interface{}{"type": "string", "description": "姓名(可选)"},
				},
				"required": []string{"year", "month", "day"},
			},
		},
		"required": []string{"person1", "person2"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		p1 := params["person1"].(map[string]interface{})
		p2 := params["person2"].(map[string]interface{})

		// Get person 1 bazi
		p1Solar := calendar.NewSolarFromYmd(
			int(p1["year"].(float64)),
			int(p1["month"].(float64)),
			int(p1["day"].(float64)),
		)
		if p1Solar == nil {
			return nil, fmt.Errorf("invalid date for person 1")
		}
		p1Lunar := calendar.NewLunarFromSolar(p1Solar)
		if p1Lunar == nil {
			return nil, fmt.Errorf("invalid lunar for person 1")
		}

		// Get person 2 bazi
		p2Solar := calendar.NewSolarFromYmd(
			int(p2["year"].(float64)),
			int(p2["month"].(float64)),
			int(p2["day"].(float64)),
		)
		if p2Solar == nil {
			return nil, fmt.Errorf("invalid date for person 2")
		}
		p2Lunar := calendar.NewLunarFromSolar(p2Solar)
		if p2Lunar == nil {
			return nil, fmt.Errorf("invalid lunar for person 2")
		}

		p1Bazi := p1Lunar.GetBaZi()
		p2Bazi := p2Lunar.GetBaZi()

		p1Name := ""
		if n, ok := p1["name"].(string); ok {
			p1Name = n
		}
		p2Name := ""
		if n, ok := p2["name"].(string); ok {
			p2Name = n
		}

		// Check compatibility based on zodiac
		p1Animal := p1Lunar.GetAnimal()
		p2Animal := p2Lunar.GetAnimal()

		// Zodiac compatibility (simplified)
		compatible := []string{
			"鼠-牛", "虎-猪", "兔-狗", "龙-鸡", "蛇-猴", "马-羊",
		}
		incompatible := []string{
			"鼠-马", "牛-羊", "虎-猴", "兔-龙", "蛇-猪", "狗-鸡",
		}

		zodiacMatch := fmt.Sprintf("%s-%s", p1Animal, p2Animal)
		zodiacMatchRev := fmt.Sprintf("%s-%s", p2Animal, p1Animal)

		var zodiacResult string
		zodiacScore := 60 // base score

		for _, c := range compatible {
			if c == zodiacMatch || c == zodiacMatchRev {
				zodiacResult = "相合"
				zodiacScore += 20
				break
			}
		}
		for _, c := range incompatible {
			if c == zodiacMatch || c == zodiacMatchRev {
				zodiacResult = "相冲"
				zodiacScore -= 15
				break
			}
		}
		if zodiacResult == "" {
			zodiacResult = "中等"
		}

		// Check day Gan compatibility
		p1DayGan := p1Lunar.GetDayGan()
		p2DayGan := p2Lunar.GetDayGan()

		// Five elements compatibility
		ganWuxing := map[string]string{
			"甲": "木", "乙": "木", "丙": "火", "丁": "火",
			"戊": "土", "己": "土", "庚": "金", "辛": "金", "壬": "水", "癸": "水",
		}

		p1Wuxing := ganWuxing[p1DayGan]
		p2Wuxing := ganWuxing[p2DayGan]

		// Simple wuxing cycle: 木生火→火生土→土生金→金生水→水生木
		var wuxingResult string
		var wuxingScore int

		pairs := []struct{ a, b, r string }{
			{"木", "火", "相生"}, {"火", "土", "相生"}, {"土", "金", "相生"},
			{"金", "水", "相生"}, {"水", "木", "相生"},
			{"木", "土", "相克"}, {"土", "水", "相克"}, {"水", "火", "相克"},
			{"火", "金", "相克"}, {"金", "木", "相克"},
		}

		for _, p := range pairs {
			if (p1Wuxing == p.a && p2Wuxing == p.b) || (p1Wuxing == p.b && p2Wuxing == p.a) {
				wuxingResult = p.r
				if p.r == "相生" {
					wuxingScore = 20
				} else {
					wuxingScore = -10
				}
				break
			}
		}
		if wuxingResult == "" {
			wuxingResult = "平和"
			wuxingScore = 5
		}

		// Calculate total score
		totalScore := zodiacScore + wuxingScore
		if totalScore > 100 {
			totalScore = 100
		}

		var overall string
		if totalScore >= 85 {
			overall = "上等"
		} else if totalScore >= 70 {
			overall = "中上"
		} else if totalScore >= 55 {
			overall = "中等"
		} else {
			overall = "需注意"
		}

		return map[string]interface{}{
			"person1": map[string]interface{}{
				"name":     p1Name,
				"date":     p1Solar.ToYmd(),
				"bazi":     fmt.Sprintf("%s %s %s %s", p1Bazi[0], p1Bazi[1], p1Bazi[2], p1Bazi[3]),
				"animal":   p1Animal,
				"dayGan":   p1DayGan,
				"wuxing":   p1Wuxing,
			},
			"person2": map[string]interface{}{
				"name":     p2Name,
				"date":     p2Solar.ToYmd(),
				"bazi":     fmt.Sprintf("%s %s %s %s", p2Bazi[0], p2Bazi[1], p2Bazi[2], p2Bazi[3]),
				"animal":   p2Animal,
				"dayGan":   p2DayGan,
				"wuxing":   p2Wuxing,
			},
			"analysis": map[string]interface{}{
				"zodiac": map[string]string{
					"result": zodiacResult,
					"score":  fmt.Sprintf("%d", zodiacScore),
				},
				"wuxing": map[string]string{
					"result": wuxingResult,
					"score":  fmt.Sprintf("%d", wuxingScore),
				},
				"totalScore": totalScore,
				"overall":    overall,
			},
		}, nil
	},
}
