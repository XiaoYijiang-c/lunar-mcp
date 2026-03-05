package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// NineStarFlyingTool calculates Nine Star Flying Palace chart
var NineStarFlyingTool = &Tool{
	Name:        "nine_star_flying",
	Description: "九星飞宫，玄空飞星风水",
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

		// Get year, month, day nine stars
		_ = lunar.GetYearNineStar()
		_ = lunar.GetMonthNineStar()
		_ = lunar.GetDayNineStar()

		// Nine star names
		nineStars := []string{"一白", "二黑", "三碧", "四绿", "五黄", "六白", "七赤", "八白", "九紫"}

		// Get index (0-8)
		getStarIndex := func(ns interface{}) int {
			return 0
		}

		// For now, use string names
		_ = getStarIndex

		// Nine star meanings
		starMeanings := map[string]string{
			"一白": "贪狼星 - 桃花运、财运、文昌",
			"二黑": "巨门星 - 病符、阴煞、困难",
			"三碧": "禄存星 - 争斗、是非、破财",
			"四绿": "文曲星 - 学业、功名、艺术",
			"五黄": "廉贞星 - 灾祸、动荡、疾病",
			"六白": "武曲星 - 事业、财运、权力",
			"七赤": "破军星 - 变化、损耗、口舌",
			"八白": "左辅星 - 财帛、贵人、稳固",
			"九紫": "右弼星 - 喜事、姻缘、远行",
		}

		// Flying palace calculation (simplified)
		// Calculate the star at each position (9 grid)
		yearNum := year - 2000
		yearStar := (11 - (yearNum % 9)) % 9

		// Get positions
		positions := []string{"北方", "东北", "东方", "东南", "中央", "西南", "西方", "西北", "南方"}

		// Build flying chart
		var chart []map[string]interface{}

		// Current star at each position
		for i := 0; i < 9; i++ {
			starIdx := (yearStar + i) % 9
			chart = append(chart, map[string]interface{}{
				"position": positions[i],
				"star":     nineStars[starIdx],
				"meaning":  starMeanings[nineStars[starIdx]],
			})
		}

		// Determine auspicious directions
		var auspicious []string
		var inauspicious []string

		for _, c := range chart {
			star := c["star"].(string)
			if star == "一白" || star == "八白" || star == "九紫" || star == "六白" {
				auspicious = append(auspicious, c["position"].(string))
			}
			if star == "二黑" || star == "五黄" || star == "三碧" {
				inauspicious = append(inauspicious, c["position"].(string))
			}
		}

		// Get day info
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
			"nineStar": map[string]interface{}{
				"yearStar":  nineStars[yearStar],
				"yearStarMeaning": starMeanings[nineStars[yearStar]],
			},
			"flyingPalace": chart,
			"auspiciousDirections":  auspicious,
			"inauspiciousDirections": inauspicious,
			"dailyFortune": map[string]interface{}{
				"yi": yiList,
				"ji": jiList,
			},
			"advice": map[string]string{
				"good":  "宜在一白、八白、九紫、六白方位活动",
				"avoid": "忌在二黑、五黄、三碧方位长时间停留",
			},
		}, nil
	},
}
