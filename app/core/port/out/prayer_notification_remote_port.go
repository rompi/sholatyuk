package port_out

import (
	"context"

	"github.com/rompi/sholatyuk/app/core/dto"
)

type PrayerNotificationRemotePort interface {
	SendNotification(ctx context.Context, notification *dto.PrayerNotification) error
}
