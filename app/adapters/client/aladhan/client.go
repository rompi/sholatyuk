package aladhan

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rompi/sholatyuk/app/core/dto"
)

// Implement the PrayerTimesService interface
type AladhanClient struct {
	baseURL string
}

// Constructor for AladhanClient
func NewAladhanClient(baseURL string) *AladhanClient {
	return &AladhanClient{baseURL: baseURL}
}

// Function to fetch prayer times
func (client *AladhanClient) GetPrayerTimes(ctx context.Context, date string) (*dto.AdhanTimes, error) {
	url := fmt.Sprintf("%s/timingsByCity/%s?city=%s&country=%s", client.baseURL, date, ctx.Value("city"), ctx.Value("country"))
	fmt.Println("URL: ", url)
	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prayer times: %w", err)
	}
	defer resp.Body.Close()

	// Decode JSON response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error1: ", resp.StatusCode)
		fmt.Println("Error1: ", resp.Body)
		fmt.Println("Error1: ", err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Return only the prayer times
	return response.Data.ConvertToAdhanTimes(), nil
}
