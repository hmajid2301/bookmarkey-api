package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}
		record := models.NewRecord(collection)
		record.Set("id", "-1")
		record.Set("name", "Unsorted Bookmarks")
		return dao.SaveRecord(record)

	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		record, err := dao.FindRecordById("pejwlxcamufi2z9", "-1")
		if err != nil {
			return err
		}

		return dao.DeleteRecord(record)
	})
}
