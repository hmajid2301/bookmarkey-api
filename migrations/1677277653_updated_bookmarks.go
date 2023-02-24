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
		edit_favourite := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kmbgnjco",
			"name": "favourite",
			"type": "bool",
			"required": false,
			"unique": false,
			"options": {}
		}`), edit_favourite)
		collection.Schema.AddField(edit_favourite)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		// update
		edit_favourite := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kmbgnjco",
			"name": "favourite",
			"type": "bool",
			"required": true,
			"unique": false,
			"options": {}
		}`), edit_favourite)
		collection.Schema.AddField(edit_favourite)

		return dao.SaveCollection(collection)
	})
}
