package tools

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// AuspiciousTimeTool finds auspicious times for the day
var AuspiciousTimeTool = &Tool{
	Name:        "auspicious_time",
	Description: "查询每日的吉时",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"year":  map[string]interface{}{"type": "integer", "description": "公历年份"},
			"month": map[string]interface{}{"type": "integer", "description": "公历月份"},
			"day":   map[string]interface{}{"type": "integer", "description": "公历日期"},
			"purpose": map[string]interface{}{"type": "string", "description": "用途: 嫁娶/搬家/开业/动土/出行 (可选)"},
		},
		"required": []string{"year", "month", "day"},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		year := int(params["year"].(float64))
		month := int(params["month"].(float64))
		day := int(params["day"].(float64))
		purpose := ""
		if v, ok := params["purpose"].(string); ok {
			purpose = v
		}

		solar := calendar.NewSolarFromYmd(year, month, day)
		if solar == nil {
			return nil, fmt.Errorf("invalid date")
		}

		// Check each of the 12 two-hour periods
		var results []map[string]interface{}

		// Two-hour periods (时辰)
		timePeriods := []struct {
			start int
			name  string
		}{
			{23, "子时"}, {1, "丑时"}, {3, "寅时"}, {5, "卯时"},
			{7, "辰时"}, {9, "巳时"}, {11, "午时"}, {13, "未时"},
			{15, "申时"}, {17, "酉时"}, {19, "戌时"}, {21, "亥时"},
		}

		for _, tp := range timePeriods {
			// Create solar with this hour
			s := calendar.NewSolar(year, month, day, tp.start, 0, 0)
			if s == nil {
				continue
			}

			lunar := calendar.NewLunarFromSolar(s)
			if lunar == nil {
				continue
			}

			// Get yi (auspicious) and ji (inauspicious) for this time
			yi := lunar.GetTimeYi()
			ji := lunar.GetTimeJi()

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

			// Determine if auspicious
			isAuspicious := len(yiList) > len(jiList)
			reason := ""

			// If purpose specified, check if matches
			if purpose != "" {
				for _, y := range yiList {
					if contains(y, purpose) {
						isAuspicious = true
						reason = y
						break
					}
				}
			}

			if isAuspicious || purpose == "" {
				results = append(results, map[string]interface{}{
					"time":       fmt.Sprintf("%s (%02d:00-%02d:00)", tp.name, tp.start, (tp.start+2)%24),
					"hour":       tp.start,
					"auspicious": yiList,
					"reason":     reason,
				})
			}
		}

		return map[string]interface{}{
			"date":     fmt.Sprintf("%d-%02d-%02d", year, month, day),
			"purpose":  purpose,
			"results":  results,
			"total":    len(results),
		}, nil
	},
}

// contains checks if s contains substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
