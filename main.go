package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	"os"
)

var URL = "https://api.telegram.org/bot"
var TELEGRAM_TOKEN = "5991269134:AAHn1LB98eG2l1eBm4SXgo_o_PAEr_4pkOo"

func main() {
	// create bot and enable debugging info
	bot, err := telego.NewBot(TELEGRAM_TOKEN, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set up a webhook on Telegram side
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL: "https://example.com/bot" + bot.Token(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Weebhook Info %+v\n", info)

	// Get updates channel from webhook
	updates, _ := bot.UpdatesViaWebhook("/bot" + bot.Token())

	//Start server for receiving request from the Telegram
	go func() {
		_ = bot.StartWebhook("localhost:443")
	}()

	// stop reviving updates from channel and shutdown webhook server
	defer func() {
		_ = bot.StopWebhook()
	}()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update %+v\n", update)
	}

}
