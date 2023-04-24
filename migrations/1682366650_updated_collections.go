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

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		// add
		new_shareable_url := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gqiaynf4",
			"name": "shareable_url",
			"type": "url",
			"required": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": []
			}
		}`), new_shareable_url)
		collection.Schema.AddField(new_shareable_url)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("gqiaynf4")

		return dao.SaveCollection(collection)
	})
}
