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
		edit_collection := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "sabinmy4",
			"name": "collection",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "pejwlxcamufi2z9",
				"cascadeDelete": false,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_collection)
		collection.Schema.AddField(edit_collection)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		// update
		edit_collection := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "sabinmy4",
			"name": "collection",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "pejwlxcamufi2z9",
				"cascadeDelete": false,
				"maxSelect": null,
				"displayFields": null
			}
		}`), edit_collection)
		collection.Schema.AddField(edit_collection)

		return dao.SaveCollection(collection)
	})
}
