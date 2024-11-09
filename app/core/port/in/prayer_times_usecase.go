package port_in

import (
	"context"

	"github.com/rompi/sholatyuk/app/core/dto"
	port_out "github.com/rompi/sholatyuk/app/core/port/out"
)

type PrayerTimesUsecase interface {
	// Get prayer times for a specific city and country
	GetPrayerTimes(ctx context.Context, date string) (*dto.AdhanTimes, error)

	NotifyClient(ctx context.Context, notif *dto.PrayerNotification) error
	SetPrayerNotificationPort(prayerNotificationPort port_out.PrayerNotificationRemotePort)
}
