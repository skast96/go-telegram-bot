package main

import (
	"HeinzBotGoEdition/bot"
	"HeinzBotGoEdition/bot/modules/kachelmann"
	"HeinzBotGoEdition/bot/modules/reddit"
)

func main() {
	registerBots()
	bot.Bot().Start()
}

func registerBots() {
	reddit.RegisterReddit()
	kachelmann.RegisterForecast()
}
