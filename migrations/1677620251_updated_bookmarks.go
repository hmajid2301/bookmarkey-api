package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		// update
		edit_bookmark_metadata := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "xst7gj3i",
			"name": "bookmark_metadata",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "elul616p7fivid4",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), edit_bookmark_metadata)
		collection.Schema.AddField(edit_bookmark_metadata)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		// update
		edit_bookmark_metadata := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "xst7gj3i",
			"name": "bookmark_metadata",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "elul616p7fivid4",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), edit_bookmark_metadata)
		collection.Schema.AddField(edit_bookmark_metadata)

		return dao.SaveCollection(collection)
	})
}
