package services

import (
	"context"
	"fmt"

	"github.com/rompi/sholatyuk/app/core/dto"
	port_in "github.com/rompi/sholatyuk/app/core/port/in"
	port_out "github.com/rompi/sholatyuk/app/core/port/out"
)

type PrayerTimesService struct {
	remotePort             port_out.PrayerTimesRemotePort
	prayerNotificationPort port_out.PrayerNotificationRemotePort
}

func NewPrayerTimesService(remotePort port_out.PrayerTimesRemotePort) port_in.PrayerTimesUsecase {
	return &PrayerTimesService{remotePort: remotePort}
}

func (s *PrayerTimesService) SetPrayerNotificationPort(prayerNotificationPort port_out.PrayerNotificationRemotePort) {
	s.prayerNotificationPort = prayerNotificationPort
}

func (s *PrayerTimesService) GetPrayerTimes(ctx context.Context, date string) (*dto.AdhanTimes, error) {
	// Fetch prayer times from remote port
	fmt.Printf("Fetching prayer times for date: %s\n", date)
	prayerTimes, err := s.remotePort.GetPrayerTimes(ctx, date)
	if err != nil {
		return nil, err
	}

	return prayerTimes, nil
}

func (s *PrayerTimesService) NotifyClient(ctx context.Context, notif *dto.PrayerNotification) error {
	// Notify the client
	s.prayerNotificationPort.SendNotification(ctx, notif)
	return nil
}
