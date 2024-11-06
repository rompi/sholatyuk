package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rompi/sholatyuk/app/adapters/client/aladhan"
	"github.com/rompi/sholatyuk/app/core/dto"
	"github.com/rompi/sholatyuk/app/core/services"
)

func main() {
	// GetScheduleByCityAndDate()
	fmt.Println("Starting sholatyuk")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "city", "Jakarta")
	ctx = context.WithValue(ctx, "country", "ID")
	apiClient := aladhan.NewAladhanClient("http://api.aladhan.com/v1")
	prayerTimesService := services.NewPrayerTimesService(apiClient)
	prayerTimes, err := prayerTimesService.GetPrayerTimes(ctx, time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println("Prayer Times:")
	fmt.Printf("Date    : %s\n", prayerTimes.Date)
	fmt.Printf("Fajr    : %s\n", prayerTimes.Fajr)
	fmt.Printf("Dhuhr   : %s\n", prayerTimes.Dhuhr)
	fmt.Printf("Asr     : %s\n", prayerTimes.Asr)
	fmt.Printf("Maghrib : %s\n", prayerTimes.Maghrib)
	fmt.Printf("Isha    : %s\n", prayerTimes.Isha)
	go checkPrayerTime(prayerTimes)

	// Keep the main function running
	select {}

}

// Function to check current time against prayer times
func checkPrayerTime(adhanTimes *dto.AdhanTimes) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		// Get current time in "HH:MM" format
		currentTime := time.Now().Format("15:04")

		// Compare with each Adhan time
		switch currentTime {
		case adhanTimes.Fajr:
			log.Println("It's time for Fajr prayer!")
		case adhanTimes.Dhuhr:
			log.Println("It's time for Dhuhr prayer!")
		case adhanTimes.Asr:
			log.Println("It's time for Asr prayer!")
		case adhanTimes.Maghrib:
			log.Println("It's time for Maghrib prayer!")
		case adhanTimes.Isha:
			log.Println("It's time for Isha prayer!")
		}
	}
}
