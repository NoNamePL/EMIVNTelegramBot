package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
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
	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {

		if update.Message != nil {
			//retrive chat ID
			chatID := update.Message.Chat.ID

			// call method sendmessage
			// Send a message to sender with the same text(echo bot)
			sentMessage, _ := bot.SendMessage(
				tu.Message(
					tu.ID(chatID),
					update.Message.Text,
				),
			)
			fmt.Printf("Setn Message: %v\n", sentMessage)

		}
	}
}
