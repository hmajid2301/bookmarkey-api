package bookmarks

import (
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Bookmark)(nil)
var _ models.Model = (*BookmarkMetaData)(nil)

// Bookmark models the data in the `bookmarks` collection.
type Bookmark struct {
	models.BaseModel

	Collection       string `db:"collection" json:"collection"`
	Favourite        bool   `db:"favourite" json:"favourite"`
	CustomOrder      int    `db:"custom_order" json:"custom_order"`
	BookmarkMetadata string `db:"bookmark_metadata" json:"bookmark_metadata"`
	User             string `db:"user" json:"user"`
}

// TableName returns the name of the model in PocketBase.
func (m *Bookmark) TableName() string {
	return "bookmarks"
}

// BookmarkMetaData is model the represents data related to URL
type BookmarkMetaData struct {
	models.BaseModel

	URL         string `db:"url" json:"url"`
	Description string `db:"description" json:"description"`
	Title       string `db:"title" json:"title"`
	Image       string `db:"image" json:"image"`
}

// TableName returns the name of the model in PocketBase.
func (m *BookmarkMetaData) TableName() string {
	return "bookmarks_metadata"
}
