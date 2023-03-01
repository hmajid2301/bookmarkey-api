package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("elul616p7fivid4")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@collection.bookmarks.collection.user.id = @request.auth.id && @collection.bookmarks.bookmark_metadata.id = id")

		collection.ViewRule = types.Pointer("@collection.bookmarks.collection.user.id = @request.auth.id && @collection.bookmarks.bookmark_metadata.id = id")

		collection.UpdateRule = types.Pointer("@collection.bookmarks.collection.user.id = @request.auth.id && @collection.bookmarks.bookmark_metadata.id = id")

		collection.DeleteRule = types.Pointer("@collection.bookmarks.collection.user.id = @request.auth.id && @collection.bookmarks.bookmark_metadata.id = id")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("elul616p7fivid4")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		collection.ViewRule = nil

		collection.UpdateRule = nil

		collection.DeleteRule = nil

		return dao.SaveCollection(collection)
	})
}
