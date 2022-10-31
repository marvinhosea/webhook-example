package app

import (
	"github.com/marvinhosea/webhook-server/config"
	"github.com/marvinhosea/webhook-server/internal"
	"github.com/marvinhosea/webhook-server/internal/repository"
	"github.com/marvinhosea/webhook-server/internal/services"
	"github.com/marvinhosea/webhook-server/internal/webhook"
)

type App struct {
	config  *config.AppConfig
	webhook webhook.Webhook
	storage internal.Storage
}

func initialize() (*App, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	db, err := internal.New(cfg.SqlDb)
	if err != nil {
		return nil, err
	}

	wb := webhook.NewClient()
	return &App{config: cfg, webhook: wb, storage: db}, nil
}

func Start() error {
	app, err := initialize()
	if err != nil {
		return err
	}
	webhookService := services.NewDefaultWebhookService(
		repository.NewDefaultEventDbRepo(app.storage),
		app.webhook)
	err = webhookService.Dispatch()
	if err != nil {
		return err
	}
	return nil
}
