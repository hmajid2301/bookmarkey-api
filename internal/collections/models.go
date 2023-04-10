// Package collections provides ways to interact with the collections, such as checking who owns a collection
package collections

import (
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Collection)(nil)

// Collection models the data in the `collections` collection.
type Collection struct {
	models.BaseModel
	Name        string `db:"name" json:"name"`
	User        string `db:"user" json:"user"`
	Group       string `db:"group" json:"group"`
	CustomOrder int    `db:"custom_order" json:"custom_order"`
}

// TableName returns the name of the model in PocketBase.
func (m *Collection) TableName() string {
	return "collections"
}
