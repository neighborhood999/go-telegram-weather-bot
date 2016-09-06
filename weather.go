package main

import (
	"errors"
	"sort"
)

const (
	tornado                 = "🌪"
	tropicalStorm           = "💨"
	hurricane               = "💨"
	severeThunderstorms     = "⛈"
	thunderstorms           = "⛈"
	mixedRainAndSnow        = "🌨"
	mixedRainAndSleet       = "🌧"
	mixedSnowAndSleet       = "🌨"
	freezingDrizzle         = "🌧"
	drizzle                 = "🌧"
	freezingRain            = "🌧"
	showers                 = "🌧"
	showerss                = "🌧"
	snowFlurries            = "🌨"
	lightSnowShowers        = "🌨"
	blowingSnow             = "🌨"
	snow                    = "🌨"
	hail                    = "🌨"
	sleet                   = "🌧"
	dust                    = "🌫"
	foggy                   = "🌫"
	haze                    = "🌫"
	smoky                   = "🌫"
	blustery                = "💨"
	windy                   = "💨"
	cody                    = "❄️"
	cloudy                  = "☁️"
	mostlyCloudyNight       = "☁️"
	mostlyCloudyDay         = "🌥"
	partlyCloudyNight       = "☁️"
	partlyCloudyDay         = "⛅️"
	clearNight              = "🌙"
	sunny                   = "☀️"
	fairNight               = "🌙"
	fairDay                 = "🌤"
	mixedRainAndHail        = "🌧"
	hot                     = "☀️"
	isolatedThunderstorms   = "⛈"
	scatteredThunderstorms  = "⛈"
	scatteredThunderstormss = "⛈"
	scatteredShowers        = "🌧"
	heavySnow               = "🌨"
	scatteredSnowShowers    = "🌨"
	heavySnows              = "🌨"
	partlyCloudy            = "⛅️"
	thundershowers          = "⛈"
	snowShowers             = "🌨"
	isolatedThundershowers  = "⛈"
	notAvailable            = "🈚️"
)

// Compass is direction and degree
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

// GuessDirection just a dumb method for get direction
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

// CheckWindDirection is return wind direction string
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

	emojiMap[0] = tornado
	emojiMap[1] = tropicalStorm
	emojiMap[2] = hurricane
	emojiMap[3] = severeThunderstorms
	emojiMap[4] = thunderstorms
	emojiMap[5] = mixedRainAndSnow
	emojiMap[6] = mixedRainAndSleet
	emojiMap[7] = mixedSnowAndSleet
	emojiMap[8] = freezingDrizzle
	emojiMap[9] = drizzle
	emojiMap[10] = freezingRain
	emojiMap[11] = showers
	emojiMap[12] = showerss
	emojiMap[13] = snowFlurries
	emojiMap[14] = lightSnowShowers
	emojiMap[15] = blowingSnow
	emojiMap[16] = snow
	emojiMap[17] = hail
	emojiMap[18] = sleet
	emojiMap[19] = dust
	emojiMap[20] = foggy
	emojiMap[21] = haze
	emojiMap[22] = smoky
	emojiMap[23] = blustery
	emojiMap[24] = windy
	emojiMap[25] = cody
	emojiMap[26] = cloudy
	emojiMap[27] = mostlyCloudyNight
	emojiMap[28] = mostlyCloudyDay
	emojiMap[29] = partlyCloudyNight
	emojiMap[30] = partlyCloudyDay
	emojiMap[31] = clearNight
	emojiMap[32] = sunny
	emojiMap[33] = fairNight
	emojiMap[34] = fairDay
	emojiMap[35] = mixedRainAndHail
	emojiMap[36] = hot
	emojiMap[37] = isolatedThunderstorms
	emojiMap[38] = scatteredThunderstorms
	emojiMap[39] = scatteredThunderstormss
	emojiMap[40] = scatteredShowers
	emojiMap[41] = heavySnow
	emojiMap[42] = scatteredSnowShowers
	emojiMap[43] = heavySnows
	emojiMap[44] = partlyCloudy
	emojiMap[45] = thundershowers
	emojiMap[46] = snowShowers
	emojiMap[47] = isolatedThundershowers
	emojiMap[3200] = notAvailable

	for k := range emojiMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)

	if emojiMap[code] == "" {
		return "", errors.New("Can't not find emojiMap.")
	}

	return emojiMap[code], nil
}
