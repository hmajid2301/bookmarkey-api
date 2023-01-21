// +gocover:ignore:file ignore this file!
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

		collection, err := dao.FindCollectionByNameOrId("collections")
		if err != nil {
			return err
		}

		// add
		new_group := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "z3kpxzdo",
			"name": "group",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "xowwnq4hswfdsci",
				"cascadeDelete": false
			}
		}`), new_group)
		collection.Schema.AddField(new_group)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("collections")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("z3kpxzdo")

		return dao.SaveCollection(collection)
	})
}
