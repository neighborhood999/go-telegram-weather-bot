package main

import (
	"errors"
	"sort"
)

func Compass() (map[float64]string, []float64) {
	var sortedKeys []float64
	direction := make(map[float64]string, 9)

	direction[0] = "無"
	direction[45] = "東北"
	direction[90] = "東"
	direction[135] = "東南"
	direction[180] = "南"
	direction[225] = "西南"
	direction[270] = "西"
	direction[315] = "西北"
	direction[360] = "北"

	for k := range direction {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Float64s(sortedKeys)

	return direction, sortedKeys
}

func GuessDirection(degree float64) float64 {
	var slice []float64
	_, sortedKeys := Compass()

	for _, k := range sortedKeys {
		if k > degree {
			slice = append(slice, k)
		}
	}

	return slice[0]
}

func CheckWindDirection(degree float64) string {
	var result string
	guessDirection := GuessDirection(degree)
	direction, _ := Compass()

	if direction[degree] != "" {
		result = direction[degree] + "風"
	} else if (guessDirection - 45) < 45 {
		result = direction[guessDirection] + "偏北風"
	} else {
		result = direction[guessDirection] + "偏" + direction[guessDirection-45] + "風"
	}

	return result
}

func weatherEmoji(code int) (string, error) {
	emojiMap := make(map[int]string, 48)
	sortedKeys := make([]int, 48)

	emojiMap[0] = "🌪"
	emojiMap[1] = "💨"
	emojiMap[2] = "💨"
	emojiMap[3] = "⛈"
	emojiMap[4] = "⛈"
	emojiMap[5] = "🌨"
	emojiMap[6] = "🌧"
	emojiMap[7] = "🌨"
	emojiMap[8] = "🌧"
	emojiMap[9] = "🌧"
	emojiMap[10] = "🌧"
	emojiMap[11] = "🌧"
	emojiMap[12] = "🌧"
	emojiMap[13] = "🌨"
	emojiMap[14] = "🌨"
	emojiMap[15] = "🌨"
	emojiMap[16] = "🌨"
	emojiMap[17] = "🌨"
	emojiMap[18] = "🌧"
	emojiMap[19] = "🌫"
	emojiMap[20] = "🌫"
	emojiMap[21] = "🌫"
	emojiMap[22] = "🌫"
	emojiMap[23] = "💨"
	emojiMap[24] = "💨"
	emojiMap[25] = "❄️"
	emojiMap[26] = "☁️"
	emojiMap[27] = "☁️"
	emojiMap[28] = "🌥"
	emojiMap[29] = "☁️"
	emojiMap[30] = "⛅️"
	emojiMap[31] = "🌙"
	emojiMap[32] = "☀️"
	emojiMap[33] = "🌙"
	emojiMap[34] = "🌤"
	emojiMap[35] = "🌧"
	emojiMap[36] = "☀️"
	emojiMap[37] = "⛈"
	emojiMap[38] = "⛈"
	emojiMap[39] = "⛈"
	emojiMap[40] = "🌧"
	emojiMap[41] = "🌨"
	emojiMap[42] = "🌨"
	emojiMap[43] = "🌨"
	emojiMap[44] = "⛅️"
	emojiMap[45] = "⛈"
	emojiMap[46] = "🌨"
	emojiMap[47] = "⛈"
	emojiMap[3200] = "🈚️"

	for k := range emojiMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)

	if emojiMap[code] == "" {
		return "", errors.New("Can't not find emojiMap.")
	}

	return emojiMap[code], nil
}
