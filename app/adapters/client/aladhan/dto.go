package aladhan

import "github.com/rompi/sholatyuk/app/core/dto"

// Define structs for parsing the response
type PrayerTimes struct {
	Fajr    string `json:"Fajr"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type Date struct {
	Readable string `json:"readable"`
}

type Data struct {
	Timings PrayerTimes `json:"timings"`
	Date    Date        `json:"date"`
}

type Response struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

// Custom AdhanTimes struct
type AdhanTimes struct {
	Date    string `json:"date"`
	Fajr    string `json:"fajr"`
	Dhuhr   string `json:"dhuhr"`
	Asr     string `json:"asr"`
	Maghrib string `json:"maghrib"`
	Isha    string `json:"isha"`
}

// Convert from Data struct to AdhanTimes
func (data *Data) ConvertToAdhanTimes() *dto.AdhanTimes {
	return &dto.AdhanTimes{
		Date:    data.Date.Readable,
		Fajr:    data.Timings.Fajr,
		Dhuhr:   data.Timings.Dhuhr,
		Asr:     data.Timings.Asr,
		Maghrib: data.Timings.Maghrib,
		Isha:    data.Timings.Isha,
	}
}
