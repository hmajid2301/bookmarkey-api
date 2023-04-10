package collections

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
)

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
func (s SQLiteStore) GetByID(id string) (*Collection, error) {

	collection := &Collection{}
	err := s.client.ModelQuery(&Collection{}).AndWhere(dbx.HashExp{"id": id}).Limit(1).One(collection)
	if err != nil {
		return nil, err
	}

	return collection, nil
}
