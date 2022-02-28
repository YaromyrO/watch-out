package main

import (
	"github.com/YaromyrO/watch-out/bot"
)

const apiToken = "5068376274:AAHsDaRFMzy7kzfLabf4tkmtGsu_sEwppMk"

func main() {
	bot := bot.NewBot(apiToken, 60)
	bot.HandleRequests()
}
