package sqlite

import (
	"encoding/json"
	"github.com/marvinhosea/webhook-server/internal/models"
	"os"
)

func (s *Sql) migration() error {
	data, err := os.ReadFile("./data/example_events.json")
	if err != nil {
		return err
	}

	var events []models.Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		return err
	}

	for _, event := range events {
		e := event
		err := s.CreateEvent(&e)
		if err != nil {
			return err
		}
	}

	return nil
}
