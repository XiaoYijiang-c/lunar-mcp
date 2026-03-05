package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// PengzuBaijiTool returns complete Pengzu Baiji (彭祖百忌)
var PengzuBaijiTool = &Tool{
	Name:        "pengzu_baiji",
	Description: "彭祖百忌完整版",
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

		// Get pengzu
		_ = lunar.GetPengZuGan()
		_ = lunar.GetPengZuZhi()

		// Get day gan/zhi
		dayGan := lunar.GetDayGan()
		dayZhi := lunar.GetDayZhi()

		// Get year/month gan/zhi
		yearGan := lunar.GetYearGan()
		yearZhi := lunar.GetYearZhi()
		monthGan := lunar.GetMonthGan()
		monthZhi := lunar.GetMonthZhi()

		// Common pengzu baiji sayings
		pengzuSayings := map[string]string{
			"甲": "甲不开仓，财物耗散",
			"乙": "乙不栽植，千株不长",
			"丙": "丙不修灶，必见灾殃",
			"丁": "丁不剃头，头必生疮",
			"戊": "戊不受田，田主不祥",
			"己": "己不破券，二比并亡",
			"庚": "庚不经络，织机虚张",
			"辛": "辛不合酱，主人不尝",
			"壬": "壬不泱水，更难提防",
			"癸": "癸不词讼，理弱敌强",
			"子": "子不问卜，自惹祸殃",
			"丑": "丑不冠带，主不还乡",
			"寅": "寅不祭祀，神鬼不尝",
			"卯": "卯不穿井，水泉不香",
			"辰": "辰不哭泣，必主重丧",
			"巳": "巳不远行，财物伏藏",
			"午": "午不苫盖，屋主更张",
			"未": "未不服药，毒气入肠",
			"申": "申不安床，诡祟入房",
			"酉": "酉不会客，醉坐颠狂",
			"戌": "戌不吃犬，作怪上床",
			"亥": "亥不嫁娶，必主分张",
		}

		ganBaiji := pengzuSayings[dayGan]
		zhiBaiji := pengzuSayings[dayZhi]

		// Get other restrictions
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

		return map[string]interface{}{
			"date":         solar.ToYmd(),
			"lunar":        lunar.ToFullString(),
			"day": map[string]interface{}{
				"gan": dayGan,
				"zhi": dayZhi,
			},
			"pengzuBaiji": map[string]interface{}{
				"dayGan": ganBaiji,
				"dayZhi": zhiBaiji,
			},
			"year": map[string]interface{}{
				"gan": yearGan,
				"zhi": yearZhi,
			},
			"month": map[string]interface{}{
				"gan": monthGan,
				"zhi": monthZhi,
			},
			"avoid": map[string]interface{}{
				"yi":  yiList,
				"ji":  jiList,
			},
		}, nil
	},
}
