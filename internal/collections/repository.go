package collections

import (
	"errors"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

const collectionName = "collections"

// ErrAlreadyExists is thrown if the record already exists in the database and we don't want duplicates
var ErrAlreadyExists = errors.New("record already exists in collection")

// SQLiteStore used to interact with DB
type SQLiteStore struct {
	client *daos.Dao
}

// NewStore returns a new store to interact with the database
func NewStore(app *pocketbase.PocketBase) SQLiteStore {
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

// Add adds a new collection record
func (s SQLiteStore) Add(name string, userID string) error {
	collection, err := s.client.FindCollectionByNameOrId(collectionName)
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)
	record.Set("name", name)
	record.Set("parent", nil)
	record.Set("user", userID)

	if err := s.client.SaveRecord(record); err != nil {
		return err
	}

	return nil
}

// Delete removes a collection record
func (s SQLiteStore) Delete(record *models.Record) error {
	err := s.client.DeleteRecord(record)
	if err != nil {
		return err
	}
	return nil
}
