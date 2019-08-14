package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yanzay/tbot"
)

func main() {
	token, present := os.LookupEnv("SLUG_JUICE_BOT_TOKEN")
	if present == false {
		log.Panic("SLUG_JUICE_BOT_TOKEN not set!")
	}

	webhook, present := os.LookupEnv("SLUG_JUICE_BOT_WEBHOOK")
	if present == false {
		log.Panic("SLUG_JUICE_BOT_WEBHOOK not set!")
	}

	bot := tbot.New(token, tbot.WithWebhook(webhook, ":8080"))

	c := bot.Client()

	count_by_user := make(map[string]int)

	match := "pooping|🥝🎂|🥝🍰"

	bot.HandleMessage("/count", func(m *tbot.Message) {
		var sb strings.Builder
		sb.WriteString("Number of poops:\n")
		for name, count := range count_by_user {
			sb.WriteString(fmt.Sprintf("%s: %v\n", name, count))
		}
		c.SendMessage(m.Chat.ID, sb.String())
		//c.SendMessage(m.Chat.ID, fmt.Sprintf("Number of poops: %v", count))
	})

	bot.HandleMessage(match, func(m *tbot.Message) {
		log.Printf("Increasing count for %s!", m.From.FirstName)
		count_by_user[m.From.FirstName]++
	})

	log.Printf("Starting bot...")
	err := bot.Start()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Finished!")
}
