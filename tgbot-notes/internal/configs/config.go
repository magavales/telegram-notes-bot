package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Token string
}

func Load() (*Config, error) {
	err := godotenv.Load("configs/bot_requirements.env")
	if err != nil {
		return nil, err
	}

	token := os.Getenv("BOT_TOKEN")

	// Проверка наличия обязательных переменных
	if token == "" {
		return nil, errors.New("необходимые переменные окружения отсутствуют")
	}

	return &Config{token}, nil
}
