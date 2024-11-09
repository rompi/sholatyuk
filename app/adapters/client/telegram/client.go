package telegram

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rompi/sholatyuk/app/core/dto"
)

type TelegramClient struct {
	// Add necessary fields here
	Tele *tgbotapi.BotAPI
}

func NewTelegramClient(token string) *TelegramClient {
	client := new(TelegramClient)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	client.Tele = bot

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return client
}

func (t *TelegramClient) Observe(ctx context.Context) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.Tele.GetUpdatesChan(u)

	for update := range updates {
		fmt.Printf("Update: %v\n", update.Message.Chat.ID)
		fmt.Printf("Update: %v\n", update.Message.Chat.UserName)
	}
}

func (tc *TelegramClient) SendNotification(ctx context.Context, notification *dto.PrayerNotification) error {
	// Add necessary code here
	fmt.Printf("Sending notification to Telegram\n")
	fmt.Printf("Client ID: %s\n", notification.ClientId)
	fmt.Printf("Title    : %s\n", notification.Title)
	fmt.Printf("Message  : %s\n", notification.Message)
	fmt.Printf("Platform : %s\n", notification.Platform)
	msg := tgbotapi.NewMessage(-4523142499, notification.Message)
	tc.Tele.Send(msg)
	return nil
}
