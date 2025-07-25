package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"tgbot-notes/internal/adapters"
	"tgbot-notes/internal/configs"
	"tgbot-notes/internal/models"
	"tgbot-notes/internal/services"
)

func Run() {
	var (
		err    error
		config *configs.Config
		logger *logrus.Logger
		dialog string
	)
	ctx := context.Background()
	logger = initLogger()
	logger.Infof("Application started")

	logger.Infof("Loading configuration")
	config, err = configs.Load()
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Infof("Configuration loaded")

	telegramAdapter, err := adapters.NewTelegramBot(config.Token)
	if err != nil {
		logger.Fatal(err)
	}

	telegramBot := services.NewTelegramService(telegramAdapter)
	handler := models.NewHandler()
	note := models.NewNote()

	for update := range telegramAdapter.GetUpdates() {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			err = handler.HandleCommands(ctx, telegramBot, update.Message)
			if err != nil {
				logger.Error(err)
			}
			dialog = update.Message.Command()
		} else {
			if dialog != "" {
				err = handler.HandleDialog(ctx, telegramBot, update.Message, note, &dialog)
				if err != nil {
					logger.Error(err)
				}
			}
		}

	}
}

func initLogger() *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	return l
}
