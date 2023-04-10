package bookmarks

import (
	"math"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
)

const collectionName = "bookmarks"
const metaCollectionName = "bookmarks_metadata"

// SQLiteStore used to interact with DB
type SQLiteStore struct {
	client *daos.Dao
}

// NewStore returns a new store to interact with the database
func NewStore(app core.App) SQLiteStore {
	client := app.Dao()
	return SQLiteStore{client: client}
}

// Create creates a new Bookmark and BookmarkMeta entry in the database
func (s SQLiteStore) Create(metadata *BookmarkMetaData, collectionID string, userID string) error {
	err := s.client.RunInTransaction(func(txDao *daos.Dao) error {
		if err := txDao.Save(metadata); err != nil {
			return err
		}

		bookmark := &Bookmark{
			BookmarkMetadata: metadata.Id,
			User:             userID,
			Favourite:        false,
			Collection:       collectionID,
			CustomOrder:      math.MaxInt32,
		}

		if err := txDao.Save(bookmark); err != nil {
			return err
		}

		return nil
	})
	return err
}

// GetByURL get metadata info from the databas
func (s SQLiteStore) GetByURL(url string) (*BookmarkMetaData, error) {
	bookmarkMetadata := &BookmarkMetaData{}

	err := s.client.ModelQuery(&BookmarkMetaData{}).
		AndWhere(dbx.NewExp("LOWER(url)={:url}", dbx.Params{
			"url": strings.ToLower(url),
		})).
		Limit(1).
		One(bookmarkMetadata)

	if err != nil {
		return nil, err
	}

	return bookmarkMetadata, nil
}
