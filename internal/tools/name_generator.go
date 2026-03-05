package tools

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/6tail/lunar-go/calendar"
)

// NameGeneratorTool generates names based on bazi
var NameGeneratorTool = &Tool{
	Name:        "name_generator",
	Description: "根据八字起名，分析五行喜忌并推荐名字",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":    map[string]interface{}{"type": "integer", "description": "年份"},
			"month":   map[string]interface{}{"type": "integer", "description": "月份"},
			"day":     map[string]interface{}{"type": "integer", "description": "日期"},
			"gender":  map[string]interface{}{"type": "string", "description": "性别: 男/女"},
			"surname": map[string]interface{}{"type": "string", "description": "姓氏(可选)"},
			"count":   map[string]interface{}{"type": "integer", "description": "生成数量，默认3"},
		},
		"required": []string{"year", "month", "day", "gender"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		day := int(params["day"].(float64))
		gender := params["gender"].(string)
		surname := ""
		if s, ok := params["surname"].(string); ok {
			surname = s
		}
		count := 3
		if c, ok := params["count"].(float64); ok {
			count = int(c)
		}

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

		// Determine missing/weak wuxing (for name补)
		var missingWuxing []string
		minCount := 100
		for _, c := range wuxingCount {
			if c < minCount {
				minCount = c
			}
		}
		for w, c := range wuxingCount {
			if c == minCount && minCount <= 1 {
				missingWuxing = append(missingWuxing, w)
			}
		}

		// If no clear missing, use day gan to determine
		dayGan := lunar.GetDayGan()
		var neededWuxing string

		switch dayGan {
		case "甲", "乙":
			neededWuxing = "火"
		case "丙", "丁":
			neededWuxing = "水"
		case "戊", "己":
			neededWuxing = "木"
		case "庚", "辛":
			neededWuxing = "土"
		case "壬", "癸":
			neededWuxing = "金"
		}

		// Character banks by wuxing
		charBanks := map[string][]string{
			"木": {"林", "森", "荣", "华", "英", "秀", "松", "柏", "桐", "桂", "芸", "芳", "萱", "薇", "菁", "莲"},
			"火": {"炎", "炳", "烨", "灿", "耀", "辉", "光", "明", "亮", "旭", "昊", "昌", "炅", "焱", "熠"},
			"土": {"坤", "厚", "德", "诚", "实", "岩", "岚", "峰", "磊", "宇", "安", "稳", "培", "庄"},
			"金": {"锋", "铭", "钧", "钰", "鑫", " Wayne", "锐", "镜", "镕", "铖", "镇", "铂", "玉"},
			"水": {"涛", "波", "澜", "泉", "洁", "清", "润", "涵", "汐", "瀚", "沐", "沛", "浚", "潇"},
		}

		// Also add characters good for the name based on gender
		maleChars := []string{"志", "伟", "强", "磊", "军", "杰", "涛", "明", "超", "勇", "祥", "瑞", "凯", "轩", "宇"}
		femaleChars := []string{"芳", "丽", "秀", "敏", "静", "婷", "雅", "芸", "菲", "雪", "梅", "兰", "萍", "燕", "珍"}

		// Generate names
		rand.Seed(time.Now().UnixNano())

		var names []map[string]interface{}
		used := make(map[string]bool)

		for i := 0; i < count*3 && len(names) < count; i++ {
			var char1, char2 string

			// Try to include needed wuxing
			var wx string
			if len(missingWuxing) > 0 && rand.Float32() > 0.3 {
				wx = missingWuxing[rand.Intn(len(missingWuxing))]
				charBanks[wx] = append(charBanks[wx], charBanks[neededWuxing]...)
			}

			// Select first character
			wxList := []string{"木", "火", "土", "金", "水"}
			if len(missingWuxing) > 0 {
				wxList = missingWuxing
			}
			wx = wxList[rand.Intn(len(wxList))]
			char1 = charBanks[wx][rand.Intn(len(charBanks[wx]))]

			// Second character based on gender
			if gender == "男" {
				char2 = maleChars[rand.Intn(len(maleChars))]
			} else {
				char2 = femaleChars[rand.Intn(len(femaleChars))]
			}

			fullName := surname + char1 + char2
			if used[fullName] {
				continue
			}
			used[fullName] = true

			names = append(names, map[string]interface{}{
				"name":    fullName,
				"char1":   char1,
				"char2":   char2,
				"meaning": fmt.Sprintf("补%s五行", wx),
			})
		}

		return map[string]interface{}{
			"bazi": map[string]string{
				"year":  bazi[0],
				"month": bazi[1],
				"day":   bazi[2],
				"time":  bazi[3],
			},
			"wuxing":  wuxingCount,
			"dayGan":  dayGan,
			"needed":  neededWuxing,
			"missing": missingWuxing,
			"gender":  gender,
			"surname": surname,
			"names":   names,
		}, nil
	},
}
