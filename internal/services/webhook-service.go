package services

import (
	"github.com/marvinhosea/webhook-server/internal/models"
	"github.com/marvinhosea/webhook-server/internal/repository"
	"github.com/marvinhosea/webhook-server/internal/webhook"
	"sync"
)

type WebhookService interface {
	Dispatch() error
}

type DefaultWebhookService struct {
	dispatcher webhook.Webhook
	repo       repository.EventRepo
}

func (d *DefaultWebhookService) Dispatch() error {
	getEvents, err := d.repo.GetEvents()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, event := range getEvents {
		wg.Add(1)
		webhookErr := make(chan error)
		payload := &models.Payload{
			WebhookCallbackUrl: event.CallbackUrl,
			Data: models.Data{
				Message:     event.Message,
				Status:      event.Status,
				ExtraFields: event.ExtraFields,
			},
		}

		go func(payload *models.Payload, c chan error, wg *sync.WaitGroup) {
			defer wg.Done()
			err := d.dispatcher.Send(payload)
			if err != nil {
				c <- err
			}
			c <- nil
		}(payload, webhookErr, &wg)

		err := <-webhookErr
		if err != nil {
			return err
		}
	}
	wg.Wait()
	return nil
}

func NewDefaultWebhookService(repo repository.EventRepo, dispatcher webhook.Webhook) WebhookService {
	return &DefaultWebhookService{dispatcher: dispatcher, repo: repo}
}
