package app

import (
	"context"
	"os"
	"tgbot-notes/internal/adapters"
	"tgbot-notes/internal/configs"
	"tgbot-notes/internal/models"
	"tgbot-notes/internal/models/handler"
	"tgbot-notes/internal/repository"
	"tgbot-notes/internal/services"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/sirupsen/logrus"
)

func Run() {
	var (
		err     error
		config  *configs.Config
		logger  *logrus.Logger
		command string
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

	postgres, err := repository.NewPostgres(config.URL)
	if err != nil {
		logger.Fatal(err)
	}
	telegramAdapter, err := adapters.NewTelegramBot(config.Token)
	if err != nil {
		logger.Fatal(err)
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logger.Fatal(err)
	}
	telegramBot := services.NewTelegramService(telegramAdapter, postgres)
	handler := handler.NewHandler()
	note := models.NewNote()
	dialog := models.NewDialog()
	replyKeyboard := models.NewReplyKeyboard()
	replyKeyboard.CreateKeyboardGetNoteByDate()

	for update := range telegramAdapter.GetUpdates() {
		if update.Message != nil {
			if update.Message.IsCommand() {
				err = handler.HandleCommands(ctx, telegramBot, update.Message, replyKeyboard)
				if err != nil {
					logger.Error(err)
				}
				command = update.Message.Command()
				dialog.SetState(command)
			} else {
				if command != "" {
					err = handler.HandleDialog(ctx, telegramBot, update.Message, note, &command, dialog, scheduler)
					if err != nil {
						logger.Errorf("%s", err.Error())
					}
				}

			}
		}
		if update.CallbackQuery != nil {
			err = handler.HandleCallback(ctx, telegramBot, update.CallbackQuery)
			if err != nil {
				logger.Error(err)
			}
		}

	}
}

func initLogger() *logrus.Logger {
	l := logrus.New()
	l.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.Stamp,
	}

	l.SetOutput(os.Stdout)
	return l
}
