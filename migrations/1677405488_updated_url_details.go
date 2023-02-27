package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("elul616p7fivid4")
		if err != nil {
			return err
		}

		collection.Name = "bookmark_metadata"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("elul616p7fivid4")
		if err != nil {
			return err
		}

		collection.Name = "url_details"

		return dao.SaveCollection(collection)
	})
}
