package port_out

import (
	"context"

	"github.com/rompi/sholatyuk/app/core/dto"
)

// Define an interface for fetching prayer times
type PrayerTimesRemotePort interface {
	GetPrayerTimes(ctx context.Context, date string) (*dto.AdhanTimes, error)
}
