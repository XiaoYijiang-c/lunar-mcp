package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// AuspiciousDateTool returns auspicious dates
var AuspiciousDateTool = &Tool{
	Name:        "auspicious_date",
	Description: "查询指定月份的吉日",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "公历年份"},
			"month": map[string]interface{}{"type": "integer", "description": "公历月份"},
			"type":  map[string]interface{}{"type": "string", "description": "类型: 嫁娶/搬家/开业/动土"},
		},
		"required": []string{"year", "month"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		queryType, _ := params["type"].(string)

		// Get days in month
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

			// Check yi/ji (auspicious/inauspicious)
			yi := lunar.GetDayYi()
			ji := lunar.GetDayJi()

			isAuspicious := false
			var reason string

			// Simple matching logic
			types := map[string][]string{
				"嫁娶": {"嫁娶", "纳采", "订盟"},
				"搬家": {"移徙", "入宅", "安门"},
				"开业": {"开市", "立券", "交易"},
				"动土": {"动土", "破土", "修造"},
			}

			if queryType != "" {
				checkList, ok := types[queryType]
				if ok {
					for _, t := range checkList {
						for i := yi.Front(); i != nil; i = i.Next() {
							// Convert interface{} to string
							if fmt.Sprintf("%v", i.Value) == t {
								isAuspicious = true
								reason = t
								break
							}
						}
					}
				}
			} else {
				// No type specified, check if more auspicious than inauspicious
				if yi.Len() > ji.Len() {
					isAuspicious = true
				}
			}

			if isAuspicious {
				results = append(results, map[string]interface{}{
					"date":  fmt.Sprintf("%d-%02d-%02d", year, month, day),
					"lunar": lunar.ToFullString(),
					"reason": reason,
				})
			}
		}

		return map[string]interface{}{
			"year":    year,
			"month":   month,
			"type":    queryType,
			"results": results,
		}, nil
	},
}
