package internal

import (
	"github.com/marvinhosea/webhook-server/internal/models"
	"github.com/marvinhosea/webhook-server/internal/platform/sqlite"
)

type Storage interface {
	GetEvents() ([]models.Event, error)
	UpdateEvent(event *models.Event) error
	CreateEvent(event *models.Event) error
}

func New(dbName string) (Storage, error) {
	return sqlite.NewSqlConn(dbName)
}
