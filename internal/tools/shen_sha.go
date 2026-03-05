package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// ShenShaTool returns detailed shen sha (神煞) information
var ShenShaTool = &Tool{
	Name:        "shen_sha",
	Description: "神煞查询，详细分析每日神煞",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "年份"},
			"month": map[string]interface{}{"type": "integer", "description": "月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "日期"},
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

		// Get basic info
		dayGan := lunar.GetDayGan()
		dayZhi := lunar.GetDayZhi()

		// Get positions
		xi := lunar.GetDayPositionXi()
		fu := lunar.GetDayPositionFu()
		cai := lunar.GetDayPositionCai()
		yangGui := lunar.GetDayPositionYangGui()
		yinGui := lunar.GetDayPositionYinGui()
		taiSui := lunar.GetDayPositionTaiSui()

		// Get tian shen
		tianShen := lunar.GetDayTianShen()
		tianShenType := lunar.GetDayTianShenType()
		tianShenLuck := lunar.GetDayTianShenLuck()

		// Get chong sha
		chong := lunar.GetChong()
		chongDesc := lunar.GetChongDesc()
		sha := lunar.GetSha()

		// Get yi/ji
		dayYi := lunar.GetDayYi()
		dayJi := lunar.GetDayJi()

		var yiList, jiList []string
		if dayYi != nil {
			for i := dayYi.Front(); i != nil; i = i.Next() {
				yiList = append(yiList, fmt.Sprintf("%v", i.Value))
			}
		}
		if dayJi != nil {
			for i := dayJi.Front(); i != nil; i = i.Next() {
				jiList = append(jiList, fmt.Sprintf("%v", i.Value))
			}
		}

		// Get xiu luck
		xiu := lunar.GetXiu()
		xiuLuck := lunar.GetXiuLuck()

		// Determine auspicious level
		var level string
		yiCount := len(yiList)
		jiCount := len(jiList)

		if yiCount >= 10 && jiCount <= 3 {
			level = "大吉"
		} else if yiCount >= 6 && jiCount <= 5 {
			level = "吉"
		} else if jiCount >= 10 {
			level = "凶"
		} else if jiCount > yiCount {
			level = "凶"
		} else {
			level = "平"
		}

		return map[string]interface{}{
			"date":     solar.ToYmd(),
			"lunar":    lunar.ToFullString(),
			"day": map[string]interface{}{
				"gan": dayGan,
				"zhi": dayZhi,
			},
			"level": level,
			"positions": map[string]string{
				"喜神":   xi,
				"福神":   fu,
				"财神":   cai,
				"阳贵神": yangGui,
				"阴贵神": yinGui,
				"太岁":   taiSui,
			},
			"tianShen": map[string]string{
				"name": tianShen,
				"type": tianShenType,
				"luck": tianShenLuck,
			},
			"chongSha": map[string]string{
				"chong": chong,
				"desc":  chongDesc,
				"sha":   sha,
			},
			"xiu": map[string]string{
				"name": xiu,
				"luck": xiuLuck,
			},
			"recommendation": map[string]interface{}{
				"yi": yiList,
				"ji": jiList,
			},
		}, nil
	},
}
