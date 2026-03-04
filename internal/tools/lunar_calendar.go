package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// LunarCalendarTool returns complete lunar calendar (黄历)
var LunarCalendarTool = &Tool{
	Name:        "lunar_calendar",
	Description: "获取完整的黄历信息",
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

		// Get all information
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

		// Get jie qi
		var jieQi string
		if jq := lunar.GetCurrentJieQi(); jq != nil {
			jieQi = jq.GetName()
		}

		// Get xiu
		xiu := lunar.GetXiu()
		xiuLuck := lunar.GetXiuLuck()
		xiuSong := lunar.GetXiuSong()

		// Get other info
		wuHou := lunar.GetWuHou()
		zheng := lunar.GetZheng()
		yueXiang := lunar.GetYueXiang()

		return map[string]interface{}{
			"date":     fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"lunar":    lunar.ToFullString(),
			"basic": map[string]interface{}{
				"year":     lunar.GetYear(),
				"month":    lunar.GetMonth(),
				"day":      lunar.GetDay(),
				"animal":   lunar.GetAnimal(),
				"weekday":  lunar.GetWeek(),
			},
			"bazi": map[string]string{
				"year":   lunar.GetYearInGanZhi(),
				"month":  lunar.GetMonthInGanZhi(),
				"day":    lunar.GetDayInGanZhi(),
				"time":   lunar.GetTimeInGanZhi(),
			},
			"fortune": map[string]interface{}{
				"xiPosition":  lunar.GetDayPositionXi(),
				"fuPosition":  lunar.GetDayPositionFu(),
				"caiPosition": lunar.GetDayPositionCai(),
				"tianShen":    lunar.GetDayTianShen(),
			},
			"auspicious":     yiList,
			"inauspicious":   jiList,
			"chong":          lunar.GetChong(),
			"sha":            lunar.GetSha(),
			"solarTerms":     jieQi,
			"xiu":            xiu,
			"xiuLuck":        xiuLuck,
			"xiuSong":        xiuSong,
			"wuHou":          wuHou,
			"zheng":          zheng,
			"yueXiang":       yueXiang,
		}, nil
	},
}
