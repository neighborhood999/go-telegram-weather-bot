package main

import (
	"testing"
)

func TestCompass(t *testing.T) {
	direction, sortedKeys := Compass()

	if direction[0] != "無" {
		t.Log(direction[0])
		t.Fail()
	} else if direction[45] != "東北" {
		t.Log(direction[45])
		t.Fail()
	} else if direction[90] != "東" {
		t.Log(direction[90])
		t.Fail()
	} else if direction[135] != "東南" {
		t.Log(direction[135])
		t.Fail()
	} else if direction[180] != "南" {
		t.Log(direction[180])
		t.Fail()
	} else if direction[225] != "西南" {
		t.Log(direction[225])
		t.Fail()
	} else if direction[270] != "西" {
		t.Log(direction[270])
		t.Fail()
	} else if direction[315] != "西北" {
		t.Log(direction[315])
		t.Fail()
	} else if direction[360] != "北" {
		t.Log(direction[360])
		t.Fail()
	}

	var degree float64 = 0
	for i := 0; i < len(sortedKeys); i++ {
		if sortedKeys[i] != degree {
			t.Fail()
			t.Log("sortedKey:", sortedKeys[i], "Degree:", degree)
		}
		degree = degree + 45
	}
}

func TestGuessDirection(t *testing.T) {
	guessDirection := GuessDirection(160)

	if guessDirection != 180 {
		t.Log(guessDirection)
		t.Fail()
	}
}

func TestCheckWindDirection(t *testing.T) {
	testDegreeOne := CheckWindDirection(160)
	testDegreeTwo := CheckWindDirection(23)
	testDegreeEast := CheckWindDirection(90)

	if testDegreeOne != "南偏東南風" {
		t.Log(testDegreeOne)
		t.Fail()
	}

	if testDegreeTwo != "東北偏北風" {
		t.Log(testDegreeTwo)
		t.Fail()
	}

	if testDegreeEast != "東風" {
		t.Log(testDegreeEast)
		t.Fail()
	}
}

func TestWeatherEmoji(t *testing.T) {
	code := 32

	if emoji, err := weatherEmoji(code); err != nil {
		t.Log(err)
		t.Fail()
	} else if emoji != "☀️" {
		t.Log(emoji)
		t.Fail()
	}
}
