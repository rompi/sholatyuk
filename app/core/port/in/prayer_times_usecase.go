package port_in

import "github.com/rompi/sholatyuk/app/core/dto"

type PrayerTimesUsecase interface {
	// Get prayer times for a specific city and country
	GetPrayerTimes(city, country, date string) (*dto.AdhanTimes, error)
}
