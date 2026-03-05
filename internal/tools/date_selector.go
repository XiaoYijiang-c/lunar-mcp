package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// DateSelectorTool selects optimal dates for specific purposes
var DateSelectorTool = &Tool{
	Name:        "date_selector",
	Description: "择日，根据目的选择最佳日期",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":   map[string]interface{}{"type": "integer", "description": "年份"},
			"month":  map[string]interface{}{"type": "integer", "description": "月份 (1-12)"},
			"purpose": map[string]interface{}{"type": "string", "description": "目的: 嫁娶/搬家/开业/动土/订盟/纳采/入学/出行/搬家/安门"},
		},
		"required": []string{"year", "month", "purpose"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		purpose := params["purpose"].(string)

		// Map purpose to related yi (auspicious) activities
		purposeMap := map[string][]string{
			"嫁娶": {"嫁娶", "纳采", "订盟", "会亲友", "求嗣"},
			"搬家": {"移徙", "入宅", "安门", "装修", "安床"},
			"开业": {"开市", "立券", "交易", "开光", "挂匾"},
			"动土": {"动土", "破土", "修造", "起基", "定磉"},
			"订盟": {"订盟", "纳采", "契约", "订合同"},
			"纳采": {"纳采", "问名", "嫁娶"},
			"入学": {"入学", "拜师", "求嗣"},
			"出行": {"出行", "移徙", "入宅"},
			"安门": {"安门", "修造", "动土"},
		}

		activities, ok := purposeMap[purpose]
		if !ok {
			activities = []string{purpose}
		}

		var results []map[string]interface{}

		for day := 1; day <= 31; day++ {
			solar := calendar.NewSolarFromYmd(year, month, day)
			if solar == nil {
				continue
			}

			lunar := calendar.NewLunarFromSolar(solar)
			if lunar == nil {
				continue
			}

			// Check yi activities
			yi := lunar.GetDayYi()
			if yi == nil {
				continue
			}

			var matchedActivities []string
			var allYi []string

			for i := yi.Front(); i != nil; i = i.Next() {
				yiStr := fmt.Sprintf("%v", i.Value)
				allYi = append(allYi, yiStr)
				for _, activity := range activities {
					if containsStr(yiStr, activity) {
						matchedActivities = append(matchedActivities, yiStr)
					}
				}
			}

			// Get ji (inauspicious)
			ji := lunar.GetDayJi()
			var jiCount int
			if ji != nil {
				jiCount = ji.Len()
			}

			// Score: more matched activities = better, fewer ji = better
			score := len(matchedActivities)*10 - jiCount

			if len(matchedActivities) > 0 {
				results = append(results, map[string]interface{}{
					"date":            fmt.Sprintf("%d-%02d-%02d", year, month, day),
					"lunar":           lunar.ToFullString(),
					"weekday":         solar.GetWeek(),
					"matched":         matchedActivities,
					"allAuspicious":   allYi,
					"inauspiciousCnt": jiCount,
					"score":           score,
				})
			}
		}

		// Sort by score (descending)
		for i := 0; i < len(results)-1; i++ {
			for j := i + 1; j < len(results); j++ {
				if results[i]["score"].(int) < results[j]["score"].(int) {
					results[i], results[j] = results[j], results[i]
				}
			}
		}

		return map[string]interface{}{
			"year":    year,
			"month":   month,
			"purpose": purpose,
			"results": results,
			"total":   len(results),
		}, nil
	},
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
