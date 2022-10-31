package repository

import (
	"github.com/marvinhosea/webhook-server/internal"
	"github.com/marvinhosea/webhook-server/internal/models"
)

type EventRepo interface {
	GetEvents() ([]models.Event, error)
	UpdateEvent(event *models.Event) error
}

type DefaultEventDbRepo struct {
	store internal.Storage
}

func (d *DefaultEventDbRepo) GetEvents() ([]models.Event, error) {
	return d.store.GetEvents()
}

func (d *DefaultEventDbRepo) UpdateEvent(event *models.Event) error {
	//TODO implement me
	panic("implement me")
}

func NewDefaultEventDbRepo(db internal.Storage) EventRepo {
	return &DefaultEventDbRepo{store: db}
}
