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

		collection, err := dao.FindCollectionByNameOrId("elul616p7fivid4")
		if err != nil {
			return err
		}

		// update
		edit_url := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jpw7marq",
			"name": "url",
			"type": "url",
			"required": true,
			"unique": true,
			"options": {
				"exceptDomains": [],
				"onlyDomains": []
			}
		}`), edit_url)
		collection.Schema.AddField(edit_url)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("elul616p7fivid4")
		if err != nil {
			return err
		}

		// update
		edit_url := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jpw7marq",
			"name": "url",
			"type": "url",
			"required": true,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": []
			}
		}`), edit_url)
		collection.Schema.AddField(edit_url)

		return dao.SaveCollection(collection)
	})
}
