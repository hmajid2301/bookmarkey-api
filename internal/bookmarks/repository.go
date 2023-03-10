package bookmarks

import (
	"math"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

const collectionName = "bookmarks"
const metaCollectionName = "bookmarks_metadata"

// SQLiteStore used to interact with DB
type SQLiteStore struct {
	client *daos.Dao
}

// BookmarkMetaData is model the represents data related to URL
type BookmarkMetaData struct {
	url         string
	description string
	title       string
	image       string
}

// NewStore returns a new store to interact with the database
func NewStore(app core.App) SQLiteStore {
	client := app.Dao()
	return SQLiteStore{client: client}
}

// Create creates a new Bookmark and BookmarkMeta entry in the database
func (s SQLiteStore) Create(metadata BookmarkMetaData, collectionID string, userID string) error {
	metadataCollection, err := s.client.FindCollectionByNameOrId(metaCollectionName)
	if err != nil {
		return err
	}

	bookmarkCollection, err := s.client.FindCollectionByNameOrId(collectionName)
	if err != nil {
		return err
	}

	s.client.RunInTransaction(func(txDao *daos.Dao) error {
		metadataRecord, err := txDao.FindFirstRecordByData(metaCollectionName, "url", metadata.url)
		if err != nil {
			metadataRecord = models.NewRecord(metadataCollection)
		}
		metadataRecord.Set("title", metadata.title)
		metadataRecord.Set("description", metadata.description)
		metadataRecord.Set("image", metadata.image)
		metadataRecord.Set("url", metadata.url)

		if err := txDao.SaveRecord(metadataRecord); err != nil {
			return err
		}

		bookmarkRecord := models.NewRecord(bookmarkCollection)
		bookmarkRecord.Set("bookmark_metadata", []string{metadataRecord.Id})
		bookmarkRecord.Set("user", []string{userID})
		bookmarkRecord.Set("favourite", false)
		bookmarkRecord.Set("collection", []string{collectionID})
		bookmarkRecord.Set("custom_order", math.MaxInt32)

		if err := txDao.SaveRecord(bookmarkRecord); err != nil {
			return err
		}

		return nil
	})

	return nil
}
