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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		// update
		edit_custom_order := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "m0osngzd",
			"name": "custom_order",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": 0,
				"max": null
			}
		}`), edit_custom_order)
		collection.Schema.AddField(edit_custom_order)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		// update
		edit_custom_order := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "m0osngzd",
			"name": "custom_order",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": 0,
				"max": null
			}
		}`), edit_custom_order)
		collection.Schema.AddField(edit_custom_order)

		return dao.SaveCollection(collection)
	})
}
