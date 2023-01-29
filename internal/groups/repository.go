package groups

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

const collectionName = "groups"

// SQLiteStore used to interact with DB
type SQLiteStore struct {
	client *daos.Dao
}

// NewStore returns a new store to interact with the database
func NewStore(app core.App) SQLiteStore {
	client := app.Dao()
	return SQLiteStore{client: client}
}

// GetByID returns a record by ID
func (s SQLiteStore) GetByID(id string) (*models.Record, error) {
	record, err := s.client.FindRecordById(collectionName, id)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// Delete removes a group record
func (s SQLiteStore) Delete(record *models.Record) error {
	err := s.client.DeleteRecord(record)
	if err != nil {
		return err
	}
	return nil
}
