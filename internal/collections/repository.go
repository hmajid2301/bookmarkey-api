package collections

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

const collectionName = "collections"

// SQLiteStore used to interact with DB
type SQLiteStore struct {
	client *daos.Dao
}

// NewStore returns a new store to interact with the database
func NewStore(app core.App) SQLiteStore {
	client := app.Dao()
	return SQLiteStore{client: client}
}

// Collection is the model in the database
type Collection struct {
	models.BaseModel
	Parent      string `json:"parent"`
	Name        string `json:"name"`
	User        string `json:"user"`
	Group       string `json:"group"`
	CustomOrder int    `json:"custom_order"`
}

// GetByID returns a record by ID
func (s SQLiteStore) GetByID(id string) (*Collection, error) {
	record, err := s.client.FindRecordById(collectionName, id)
	if err != nil {
		return nil, err
	}
	c := Collection{}
	// TODO: find better way to fix
	jsonRecord, _ := record.MarshalJSON()
	json.Unmarshal(jsonRecord, &c)
	return &c, nil
}
