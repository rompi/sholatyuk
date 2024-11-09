package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rompi/sholatyuk/app/adapters/client/aladhan"
	"github.com/rompi/sholatyuk/app/adapters/client/telegram"
	"github.com/rompi/sholatyuk/app/core/dto"
	port_in "github.com/rompi/sholatyuk/app/core/port/in"
	"github.com/rompi/sholatyuk/app/core/services"
	"github.com/rompi/sholatyuk/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	// GetScheduleByCityAndDate()
	fmt.Println("Starting sholatyuk")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "city", "Jakarta")
	ctx = context.WithValue(ctx, "country", "ID")
	apiClient := aladhan.NewAladhanClient(config.GlobalConfig.PrayerTimesApiBaseUrl)

	telegramClient := telegram.NewTelegramClient(config.GlobalConfig.TelegramBotToken)
	prayerTimesService := services.NewPrayerTimesService(apiClient)
	prayerTimesService.SetPrayerNotificationPort(telegramClient)
	adhanTimes, err := prayerTimesService.GetPrayerTimes(ctx, time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "Prayer Times", fmt.Sprintf("Fajr: %s\nDhuhr: %s\nAsr: %s\nMaghrib: %s\nIsha: %s\n", adhanTimes.Fajr, adhanTimes.Dhuhr, adhanTimes.Asr, adhanTimes.Maghrib, adhanTimes.Isha), "telegram"})

	fmt.Println("Prayer Times:")
	fmt.Printf("Date    : %s\n", adhanTimes.Date)
	fmt.Printf("Fajr    : %s\n", adhanTimes.Fajr)
	fmt.Printf("Dhuhr   : %s\n", adhanTimes.Dhuhr)
	fmt.Printf("Asr     : %s\n", adhanTimes.Asr)
	fmt.Printf("Maghrib : %s\n", adhanTimes.Maghrib)
	fmt.Printf("Isha    : %s\n", adhanTimes.Isha)
	go checkPrayerTime(ctx, adhanTimes, prayerTimesService)

	// Keep the main function running
	select {}

}

// Function to check current time against prayer times
func checkPrayerTime(ctx context.Context, adhanTimes *dto.AdhanTimes, prayerTimesService port_in.PrayerTimesUsecase) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		// Get current time in "HH:MM" format
		currentTime := time.Now().Format("15:04")

		if currentTime == "23:52" {
			prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "52", "jam piro iki", "telegram"})

		}

		// Compare with each Adhan time
		switch currentTime {
		case "00:01":
			adhanTimes, _ = prayerTimesService.GetPrayerTimes(ctx, time.Now().Format("2006-01-02"))

			fmt.Println("Prayer Times:")
			fmt.Printf("Date    : %s\n", adhanTimes.Date)
			fmt.Printf("Fajr    : %s\n", adhanTimes.Fajr)
			fmt.Printf("Dhuhr   : %s\n", adhanTimes.Dhuhr)
			fmt.Printf("Asr     : %s\n", adhanTimes.Asr)
			fmt.Printf("Maghrib : %s\n", adhanTimes.Maghrib)
			fmt.Printf("Isha    : %s\n", adhanTimes.Isha)
			prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "Prayer Times", fmt.Sprintf("Fajr: %s\nDhuhr: %s\nAsr: %s\nMaghrib: %s\nIsha: %s\n", adhanTimes.Fajr, adhanTimes.Dhuhr, adhanTimes.Asr, adhanTimes.Maghrib, adhanTimes.Isha), "telegram"})
		case adhanTimes.Fajr:
			prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "Fajr", "It's time for Fajr prayer!", "telegram"})
		case adhanTimes.Dhuhr:
			prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "Dhuhr", "It's time for Dhuhr prayer!", "telegram"})
		case adhanTimes.Asr:
			prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "Asr", "It's time for Asr prayer!", "telegram"})
		case adhanTimes.Maghrib:
			prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "Maghrib", "It's time for Maghrib prayer!", "telegram"})
		case adhanTimes.Isha:
			prayerTimesService.NotifyClient(ctx, &dto.PrayerNotification{"1", "Isha", "It's time for Isha prayer!", "telegram"})
		}
	}
}
