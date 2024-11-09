package config

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	AppPort     string `env:"APP_PORT"`
	AppEnv      string `env:"APP_ENV"`
	AppName     string `env:"APP_NAME"`
	AppSlugName string `env:"APP_SLUG_NAME"`
	AppVersion  string `env:"APP_VERSION"`

	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`

	PrayerTimesApiBaseUrl string `env:"PRAYER_TIMES_API_BASE_URL"`

	StartTime time.Time
}

var GlobalConfig AppConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	if err := envconfig.Process(ctx, &GlobalConfig); err != nil {
		log.Fatal(err)
	}
}
