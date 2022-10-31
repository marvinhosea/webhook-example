package sqlite

import (
	"database/sql"
	"encoding/json"
	"github.com/marvinhosea/webhook-server/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Sql struct {
	client *sql.DB
}

func (s *Sql) CreateEvent(event *models.Event) error {
	sqlQueryStatement := `INSERT INTO events(callback_url, message, status, extra_fields) VALUES (?,?,?,?)`
	prepareStatement, err := s.client.Prepare(sqlQueryStatement)
	if err != nil {
		return err
	}

	extraFields, err := json.Marshal(event.ExtraFields)
	if err != nil {
		return err
	}
	_, err = prepareStatement.Exec(event.CallbackUrl, event.Message, event.Status, string(extraFields))
	if err != nil {
		return err
	}
	return nil
}

func (s *Sql) GetEvents() ([]models.Event, error) {
	sqlQuery := `SELECT * FROM events`
	prepareStatement, err := s.client.Prepare(sqlQuery)
	if err != nil {
		return nil, err
	}
	row, err := prepareStatement.Query()
	if err != nil {
		return nil, err
	}

	var events []models.Event
	events = []models.Event{}
	defer row.Close()
	for row.Next() {
		var id int
		var callback_url, message, status, extra_fields string
		row.Scan(&id, &callback_url, &message, &status, &extra_fields)
		var extraFields []models.ExtraField
		err := json.Unmarshal([]byte(extra_fields), &extraFields)
		if err != nil {
			return nil, err
		}
		if len(extraFields) == 0 {
			extraFields = []models.ExtraField{}
		}
		event := models.Event{
			CallbackUrl: callback_url,
			Message:     message,
			Status:      status,
			ExtraFields: extraFields,
		}
		events = append(events, event)
	}
	return events, nil
}

func (s *Sql) UpdateEvent(event *models.Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *Sql) createTable() error {
	sqlQueryStatement := `CREATE TABLE events (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"callback_url" TEXT,
		"message" TEXT,
		"status" TEXT,
		"extra_fields" TEXT		
	  );`

	prepareStatement, err := s.client.Prepare(sqlQueryStatement)
	if err != nil {
		return err
	}

	_, err = prepareStatement.Exec()
	if err != nil {
		return err
	}
	log.Println("table created successfully")
	return nil
}

func initialize(dbName string) error {
	_ = os.Remove("./data/" + dbName)

	log.Println("creating sqlite db")
	file, err := os.Create("./data/" + dbName)
	defer file.Close()
	if err != nil {
		return err
	}
	log.Println("sqlite db created successfully")
	return nil
}

func NewSqlConn(dbName string) (*Sql, error) {
	err := initialize(dbName)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", "./data/"+dbName)
	if err != nil {
		return nil, err
	}

	sqlClient := &Sql{client: db}
	err = sqlClient.createTable()
	if err != nil {
		return nil, err
	}

	err = sqlClient.migration()
	if err != nil {
		return nil, err
	}

	return sqlClient, nil
}
