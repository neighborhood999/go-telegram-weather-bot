package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// CityName is search city
type CityName struct {
	Name string
}

// WeatherInfo is Bot API response information for user
type WeatherInfo struct {
	City          string
	Time          string
	Tempture      string
	Humidity      string
	Status        int
	WindSpeed     string
	WindDirection float64
	Sunrise       string
	Sunset        string
	Link          string
}

// ReadBotToken is read bot token(token.json)
func ReadBotToken(path string) (string, error) {
	var data map[string]string

	file, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return "", err
	}

	return data["token"], nil
}

// BuildURL is generate API URL
func BuildURL(param interface{}) (urlParsed string) {
	URL, _ := url.Parse("https://query.yahooapis.com/v1/public/yql")

	parameters := url.Values{}

	switch t := param.(type) {
	case tgbotapi.Location:
		latitude := strconv.FormatFloat(t.Latitude, 'E', -1, 64)
		longitude := strconv.FormatFloat(t.Longitude, 'E', -1, 64)

		parameters.Add(
			"q",
			"select * from weather.forecast where u=\"u\" AND woeid in (select woeid from geo.places(1) where text=\"("+latitude+","+longitude+")\")",
		)
	case CityName:
		parameters.Add(
			"q",
			"select * from weather.forecast where u=\"u\" AND woeid in (select woeid from geo.places(1) where text=\"("+t.Name+")\")",
		)
	}

	parameters.Add("format", "json")
	URL.RawQuery = parameters.Encode()
	urlParsed = URL.String()

	return urlParsed
}

// HTTPGet is
func HTTPGet(weatherURL string) ([]byte, error) {
	response, err := http.Get(weatherURL)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// HandleQueryResult is handle HTTPGet body
func (w *WeatherInfo) HandleQueryResult(body []byte) error {
	json, err := simplejson.NewJson(body)

	if err != nil {
		return err
	}

	city, _ := json.Get("query").Get("results").Get("channel").Get("location").Get("city").String()
	tempture, _ := json.Get("query").Get("results").Get("channel").Get("item").Get("condition").Get("temp").String()
	humidity, _ := json.Get("query").Get("results").Get("channel").Get("atmosphere").Get("humidity").String()
	status, _ := json.Get("query").Get("results").Get("channel").Get("item").Get("condition").Get("code").String()
	windSpeed, _ := json.Get("query").Get("results").Get("channel").Get("wind").Get("speed").String()
	direction, _ := json.Get("query").Get("results").Get("channel").Get("wind").Get("direction").String()
	sunrise, _ := json.Get("query").Get("results").Get("channel").Get("astronomy").Get("sunrise").String()
	sunset, _ := json.Get("query").Get("results").Get("channel").Get("astronomy").Get("sunset").String()
	link, _ := json.Get("query").Get("results").Get("channel").Get("link").String()

	if _, err := strconv.ParseFloat(direction, 64); err != nil {
		return err
	}

	if _, err := strconv.Atoi(status); err != nil {
		return err
	}

	windDirection, _ := strconv.ParseFloat(direction, 64)
	emojiStatusCode, _ := strconv.Atoi(status)

	w.City = city
	w.Time = time.Now().Format("2006-01-02 15:04:05")
	w.Tempture = tempture
	w.Humidity = humidity
	w.Status = emojiStatusCode
	w.WindSpeed = windSpeed
	w.WindDirection = windDirection
	w.Sunrise = sunrise
	w.Sunset = sunset
	w.Link = link

	return nil
}

// ResponseWeatherText is response for user's template
func (w *WeatherInfo) ResponseWeatherText(weatherInfo *WeatherInfo) string {
	emoji, _ := weatherEmoji(weatherInfo.Status)
	weatherMessage := `🚩 *` + weatherInfo.City + `*
- - - - - - - - - - - - - - - - - - - - - -
🕘 目前時間 ➡️ ` + weatherInfo.Time + `
🔰 目前溫度 ➡️ ` + weatherInfo.Tempture + `°C
💧 目前濕度 ➡️ ` + weatherInfo.Humidity + `%
🌀 天氣狀態 ➡️ ` + emoji + `
💨 目前風速 ➡️ ` + weatherInfo.WindSpeed + ` km/h
🔃 風速風向 ➡️ ` + CheckWindDirection(weatherInfo.WindDirection) + `
- - - - - - - - - - - - - - - - - - - - - -
🌅 日出時間 ➡️ ` + weatherInfo.Sunrise + `
🌄 日落時間 ➡️ ` + weatherInfo.Sunset + `

詳細資訊 🔍 [Yahoo Weather](` + weatherInfo.Link + `)
`

	return weatherMessage
}
