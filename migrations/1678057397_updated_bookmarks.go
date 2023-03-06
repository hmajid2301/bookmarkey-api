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

		// add
		new_user := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "v9ncctkk",
			"name": "user",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), new_user)
		collection.Schema.AddField(new_user)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("v9ncctkk")

		return dao.SaveCollection(collection)
	})
}
