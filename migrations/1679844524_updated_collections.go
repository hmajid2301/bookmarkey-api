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

		// remove
		collection.Schema.RemoveField("awpaxjtv")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		// add
		del_parent := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "awpaxjtv",
			"name": "parent",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "pejwlxcamufi2z9",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_parent)
		collection.Schema.AddField(del_parent)

		return dao.SaveCollection(collection)
	})
}
