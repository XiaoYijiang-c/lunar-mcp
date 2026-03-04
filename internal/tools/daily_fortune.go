package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// DailyFortuneTool returns daily fortune information
var DailyFortuneTool = &Tool{
	Name:        "daily_fortune",
	Description: "获取每日运势，包括喜神、福神、财神方位等",
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

		// Get fortune positions
		xiPos := lunar.GetDayPositionXi()
		xiDesc := lunar.GetDayPositionXiDesc()
		fuPos := lunar.GetDayPositionFu()
		fuDesc := lunar.GetDayPositionFuDesc()
		caiPos := lunar.GetDayPositionCai()
		caiDesc := lunar.GetDayPositionCaiDesc()

		// Get tian shen
		tianShen := lunar.GetDayTianShen()
		tianShenLuck := lunar.GetDayTianShenLuck()
		tianShenType := lunar.GetDayTianShenType()

		// Get yi/ji lists
		yi := lunar.GetDayYi()
		ji := lunar.GetDayJi()

		var yiList, jiList []string
		if yi != nil {
			for i := yi.Front(); i != nil; i = i.Next() {
				yiList = append(yiList, fmt.Sprintf("%v", i.Value))
			}
		}
		if ji != nil {
			for i := ji.Front(); i != nil; i = i.Next() {
				jiList = append(jiList, fmt.Sprintf("%v", i.Value))
			}
		}

		return map[string]interface{}{
			"date": fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"lunar": lunar.ToFullString(),
			"fortune": map[string]interface{}{
				"xiPosition":    xiPos,
				"xiDesc":        xiDesc,
				"fuPosition":    fuPos,
				"fuDesc":        fuDesc,
				"caiPosition":   caiPos,
				"caiDesc":      caiDesc,
				"tianShen":     tianShen,
				"tianShenLuck": tianShenLuck,
				"tianShenType": tianShenType,
			},
			"auspicious":   yiList,
			"inauspicious": jiList,
		}, nil
	},
}
