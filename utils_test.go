package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func TestReadBotToken(t *testing.T) {
	if token, err := ReadBotToken("./token.json"); err != nil {
		t.Log(err)
		t.Fail()
	} else if token == "" {
		t.Log("Invalid token")
		t.Fail()
	}

	if _, err := ReadBotToken("./testToken.json"); err != nil {
		t.Log(err)
	}

	if _, err := ReadBotToken(""); err != nil {
		t.Log(err)
	}
}

func TestBuildURL(t *testing.T) {
	location := tgbotapi.Location{Latitude: 25.047760, Longitude: 121.531850}
	expectedLocationURL := `https://query.yahooapis.com/v1/public/yql?` +
		`format=json&q=select+%2A+from+weather.forecast+where+u%3D%22u%22+` +
		`AND+woeid+in+%28select+woeid+from+geo.places%281%29+where+text%3D%22%282.504776E%2B01%2C1.2153185E%2B02%29%22%29`
	locationURL := BuildURL(location)

	if locationURL != expectedLocationURL {
		t.Log("URL Error")
		t.Fail()
	}

	city := CityName{"Taipei"}
	expectedCityURL := `https://query.yahooapis.com/v1/public/yql?` +
		`format=json&q=select+%2A+from+weather.forecast+where+u%3D%22u%22+` +
		`AND+woeid+in+%28select+woeid+from+geo.places%281%29+where+text%3D%22%28Taipei%29%22%29`
	cityURL := BuildURL(city)

	if cityURL != expectedCityURL {
		t.Log("URL Error")
		t.Fail()
	}
}

func TestResponseWeatherText(t *testing.T) {
	fakeInfo := &WeatherInfo{
		"Taipei, Taiwan",
		time.Now().Format("2006/01/02 15:04:05"),
		"87",
		"87",
		87,
		"87.87",
		87,
		"6:00",
		"6:00",
		"http://goweatherbot.example.com",
	}

	result := ResponseWeatherText(fakeInfo)

	if reflect.TypeOf(result).Kind() != reflect.String {
		t.Log("Response Fail")
		t.Fail()
	}
}