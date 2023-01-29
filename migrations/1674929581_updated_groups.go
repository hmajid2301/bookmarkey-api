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

		collection, err := dao.FindCollectionByNameOrId("xowwnq4hswfdsci")
		if err != nil {
			return err
		}

		// update
		edit_custom_order := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "fkc0oxbe",
			"name": "custom_order",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": 0,
				"max": 9007199254740991
			}
		}`), edit_custom_order)
		collection.Schema.AddField(edit_custom_order)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("xowwnq4hswfdsci")
		if err != nil {
			return err
		}

		// update
		edit_custom_order := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "fkc0oxbe",
			"name": "custom_order",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), edit_custom_order)
		collection.Schema.AddField(edit_custom_order)

		return dao.SaveCollection(collection)
	})
}
