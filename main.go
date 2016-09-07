package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Markdown is send message parse mode
const Markdown = tgbotapi.ModeMarkdown

var cityName = &CityName{}
var weatherInfo = &WeatherInfo{}
var forest = &Forest{}
var callbackConfing = &tgbotapi.CallbackConfig{}
var responseMessage string

func main() {
	token, err := ReadBotToken("./token.json")

	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil || update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "today":
				inlineButton := tgbotapi.NewInlineKeyboardButtonData("未來一週天氣", "forest")
				inlineKeyboard := []tgbotapi.InlineKeyboardButton{
					inlineButton,
				}
				inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

				e := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					responseMessage,
				)
				e.BaseEdit.ReplyMarkup = &inlineKeyboardMarkup
				e.ParseMode = Markdown

				bot.Send(e)
			case "forest":
				inlineButton := tgbotapi.NewInlineKeyboardButtonData("今日天氣狀況", "today")
				inlineKeyboard := []tgbotapi.InlineKeyboardButton{
					inlineButton,
				}
				inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

				e := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					callbackConfing.Text,
				)
				e.BaseEdit.ReplyMarkup = &inlineKeyboardMarkup
				e.ParseMode = Markdown

				bot.Send(e)
			}

			continue
		}

		command := regexp.MustCompile("/[a-z]+").FindString(update.Message.Text)

		switch command {
		case "/start":
			message := "Hello！我是 WeatherBot 😉 \n輸入 `/help` 了解如何使用 Weather Bot！"
			response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			response.ParseMode = Markdown

			bot.Send(response)
		case "/help":
			message := "輸入 `/` 有命令提示可以使用，可以透過定位查詢天氣，\n或者輸入 `/where cityName` 來尋找該地區的天氣資訊。"
			response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			response.ParseMode = Markdown

			bot.Send(response)
		case "/where":
			city := strings.Fields(update.Message.Text)
			cityName.Name = city[1]

			weatherAPIURL := BuildURL(*cityName)
			body, err := HTTPGet(weatherAPIURL)

			if err != nil {
				log.Fatal(err)
			}

			forestInfo, _ := weatherInfo.HandleQueryResult(body)
			forestResponse := forest.HandleQueryForest(forestInfo)

			responseMessage = weatherInfo.ResponseWeatherText(weatherInfo)

			callbackConfing.Text = forestResponse

			inlineButton := tgbotapi.NewInlineKeyboardButtonData("未來一週天氣", "forest")
			inlineKeyboard := []tgbotapi.InlineKeyboardButton{
				inlineButton,
			}
			inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

			response := tgbotapi.NewMessage(update.Message.Chat.ID, responseMessage)
			response.BaseChat.ReplyMarkup = inlineKeyboardMarkup
			response.ParseMode = Markdown

			bot.Send(response)
		case "/location":
			button := []tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButtonLocation("📍 取得定位"),
			}
			replyMarkup := tgbotapi.NewReplyKeyboard(button)
			replyMarkup.OneTimeKeyboard = true

			response := tgbotapi.NewMessage(update.Message.Chat.ID, "按下按鈕取得定位！")
			response.BaseChat.ReplyMarkup = replyMarkup

			bot.Send(response)
		default:
			if update.Message.Location != nil {
				weatherAPIURL := BuildURL(*update.Message.Location)
				body, err := HTTPGet(weatherAPIURL)

				if err != nil {
					log.Fatal(err)
				}

				forestInfo, _ := weatherInfo.HandleQueryResult(body)
				forestResponse := forest.HandleQueryForest(forestInfo)

				responseMessage = weatherInfo.ResponseWeatherText(weatherInfo)

				callbackConfing.Text = forestResponse

				inlineButton := tgbotapi.NewInlineKeyboardButtonData("未來一週天氣", "forest")
				inlineKeyboard := []tgbotapi.InlineKeyboardButton{
					inlineButton,
				}
				inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

				response := tgbotapi.NewMessage(update.Message.Chat.ID, responseMessage)
				response.BaseChat.ReplyMarkup = inlineKeyboardMarkup
				response.ParseMode = Markdown

				bot.Send(response)
			}
		}
	}
}
