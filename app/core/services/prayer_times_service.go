package services

import (
	"context"

	"github.com/rompi/sholatyuk/app/core/dto"
	port_out "github.com/rompi/sholatyuk/app/core/port/out"
)

type PrayerTimesService struct {
	remotePort port_out.PrayerTimesRemotePort
}

func NewPrayerTimesService(remotePort port_out.PrayerTimesRemotePort) *PrayerTimesService {
	return &PrayerTimesService{remotePort: remotePort}
}

func (s *PrayerTimesService) GetPrayerTimes(ctx context.Context, date string) (*dto.AdhanTimes, error) {
	// Fetch prayer times from remote port
	prayerTimes, err := s.remotePort.GetPrayerTimes(ctx, date)
	if err != nil {
		return nil, err
	}

	return prayerTimes, nil
}
