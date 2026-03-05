package tools

import (
	"fmt"
	"math/rand"
	"time"
)

// IChingDivinationTool performs I Ching divination
var IChingDivinationTool = &Tool{
	Name:        "iching_divination",
	Description: "易经占卜，摇卦算命",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"question": map[string]interface{}{"type": "string", "description": "所问问题(可选)"},
		},
	},
	Handler: func(params map[string]interface{}) (interface{}, error) {
		question := ""
		if q, ok := params["question"].(string); ok {
			question = q
		}

		// Use current time for divination
		now := time.Now()

		// Generate 6 yao (lines) using coins
		// 3 coins per yao, 6 yaos total
		rand.Seed(now.UnixNano())

		var yaos []int

		for i := 0; i < 6; i++ {
			// 3 coins: heads=3, tails=2
			sum := 0
			for j := 0; j < 3; j++ {
				if rand.Float32() > 0.5 {
					sum += 3 // heads
				} else {
					sum += 2 // tails
				}
			}
			// 6=old yang, 7=young yang, 8=young yin, 9=old yin
			yaos = append(yaos, sum)
		}

		// Determine hexagram
		// Convert to binary: 6/8=0 (yin), 7/9=1 (yang)
		lowerTrigram := 0
		upperTrigram := 0

		for i := 0; i < 3; i++ {
			if yaos[i] == 6 || yaos[i] == 9 {
				lowerTrigram |= (1 << i)
			}
		}
		for i := 3; i < 6; i++ {
			if yaos[i] == 6 || yaos[i] == 9 {
				upperTrigram |= (1 << (i - 3))
			}
		}

		// Trigram names
		trigrams := []string{"坤", "震", "坎", "兑", "艮", "离", "巽", "乾"}
		trigrams2 := []string{"☷", "☳", "☵", "☱", "☶", "☲", "☴", "☰"}

		lowerName := trigrams[lowerTrigram]
		upperName := trigrams[upperTrigram]
		lowerSymbol := trigrams2[lowerTrigram]
		upperSymbol := trigrams2[upperTrigram]

		// Determine hexagram name based on trigram combination
		hexagrams := map[int]string{
			1: "乾", 2: "坤", 3: "屯", 4: "蒙", 5: "需", 6: "讼", 7: "师", 8: "比",
			9: "小畜", 10: "履", 11: "泰", 12: "否", 13: "同人", 14: "大有", 15: "谦", 16: "豫",
			17: "随", 18: "蛊", 19: "临", 20: "观", 21: "噬嗑", 22: "贲", 23: "剥", 24: "复",
			25: "无妄", 26: "大畜", 27: "颐", 28: "大过", 29: "坎", 30: "离", 31: "咸", 32: "恒",
			33: "遁", 34: "大壮", 35: "晋", 36: "明夷", 37: "家人", 38: "睽", 39: "蹇", 40: "解",
			41: "损", 42: "益", 43: "夬", 44: "姤", 45: "萃", 46: "升", 47: "困", 48: "井",
			49: "革", 50: "鼎", 51: "震", 52: "艮", 53: "渐", 54: "归妹", 55: "丰", 56: "旅",
			57: "巽", 58: "兑", 59: "涣", 60: "节", 61: "中孚", 62: "小过", 63: "既济", 64: "未济",
		}

		hexagramNum := upperTrigram*8 + lowerTrigram + 1
		hexagramName := hexagrams[hexagramNum]

		// Determine if there are changing lines
		var changingLines []int
		for i, yao := range yaos {
			if yao == 6 || yao == 9 {
				changingLines = append(changingLines, i+1)
			}
		}

		// Yao meanings
		yaoMeanings := map[int]string{
			6: "老阳（变爻）", 7: "少阳", 8: "少阴", 9: "老阴（变爻）",
		}

		var yaoDetails []map[string]interface{}
		for i, yao := range yaos {
			yaoDetails = append(yaoDetails, map[string]interface{}{
				"position": i + 1,
				"value":    yao,
				"meaning":  yaoMeanings[yao],
			})
		}

		// Interpretation based on hexagram
		interpretations := map[string]string{
			"乾": "卦象为天，象征刚健有力。利于刚正不阿、有为之人。",
			"坤": "卦象为地，象征柔顺宽厚。利于稳扎稳打、渐进积累。",
			"屯": "卦象为水雷屯，象征困难重重、万物始生。",
			"蒙": "卦象为山水蒙，象征蒙昧不明、需要启蒙。",
			"需": "卦象为水天需，象征等待时机、积蓄力量。",
			"讼": "卦象为天水讼，象征争论诉讼、需谨慎行事。",
			"泰": "卦象为地天泰，象征通泰吉祥、否极泰来。",
			"否": "卦象为天地否，象征阻塞不通、需静待变化。",
		}

		interpretation := interpretations[hexagramName]
		if interpretation == "" {
			interpretation = fmt.Sprintf("第%d卦 %s，卦象复杂，需细研", hexagramNum, hexagramName)
		}

		return map[string]interface{}{
			"question":         question,
			"time":             now.Format("2006-01-02 15:04:05"),
			"hexagram": map[string]interface{}{
				"number":  hexagramNum,
				"name":    hexagramName,
				"lower":   lowerName,
				"upper":   upperName,
				"symbol":  lowerSymbol + upperSymbol,
			},
			"yaos":           yaoDetails,
			"changingLines":  changingLines,
			"interpretation": interpretation,
			"advice":         "易经占卜仅供娱乐参考，人生还需自己努力",
		}, nil
	},
}
