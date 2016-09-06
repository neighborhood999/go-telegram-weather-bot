# Telegram Weather Bot (Go)

[![Build Status](https://img.shields.io/travis/neighborhood999/go-telegram-weather-bot.svg?style=flat-square)](https://travis-ci.org/neighborhood999/go-telegram-weather-bot)
[![codecov](https://img.shields.io/codecov/c/github/neighborhood999/go-telegram-weather-bot.svg?style=flat-square)](https://codecov.io/gh/neighborhood999/go-telegram-weather-bot)
[![Go Report Card](https://goreportcard.com/badge/github.com/neighborhood999/go-telegram-weather-bot?style=flat-square)](https://goreportcard.com/report/github.com/neighborhood999/go-telegram-weather-bot)
[![telegram bot](https://img.shields.io/badge/telegram-bot-blue.svg?style=flat-square)](https://github.com/neighborhood999/go-telegram-weather-bot)

![](./screenshot/telegramBot-written-in-Go.png)

> Telegram Weather Bot written in Go.

## Usage

**First, You need have telegram account!**  

Add [@BotFather](https://telegram.me/BotFather) and send `/newbot` command to make new Bot, then you will get bot token.  

After get bot token then write token into `token.json` then running:
```sh
$ go get github.com/go-telegram-bot-api/telegram-bot-api
$ go get github.com/bitly/go-simplejson

$ go run main.go utils.go weather.go
```

## Screenshot

![telegram-weather-bot](./screenshot/tg-weather-bot.jpg)

## Test

```sh
$ go test
```

## Build

```sh
$ go build
```

## Weather API

☀️ [Yahoo Weather](https://developer.yahoo.com/weather/)

## Related

- [Telegram Bot 開發起手式](http://neighborhood999.github.io/2016/07/19/Develop-telegram-bot/)
- [Telegram Weather Bot _(Nodejs)_](https://github.com/neighborhood999/telegram-weather-bot)

## LICENSE

[Peng Jie](https://github.com/neighborhood999) © MIT
