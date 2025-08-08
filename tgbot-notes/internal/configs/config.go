package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Token string
	URL   string
}

func Load() (*Config, error) {
	var err error

	err = godotenv.Load("configs/bot_requirements.env")
	if err != nil {
		return nil, err
	}

	err = godotenv.Load("configs/postgres.env")
	if err != nil {
		return nil, err
	}

	token := os.Getenv("BOT_TOKEN")
	url := os.Getenv("URL")

	// Проверка наличия обязательных переменных
	if token == "" {
		return nil, errors.New("необходимые переменные окружения отсутствуют")
	}

	conf := &Config{
		Token: token,
		URL:   url,
	}

	return conf, nil
}
