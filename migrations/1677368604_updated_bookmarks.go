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

		// remove
		collection.Schema.RemoveField("_pb_users_auth_")

		// remove
		collection.Schema.RemoveField("lgfvbvbz")

		// add
		new_url := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "xst7gj3i",
			"name": "url",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "elul616p7fivid4",
				"cascadeDelete": false,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), new_url)
		collection.Schema.AddField(new_url)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		// add
		del_user := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "_pb_users_auth_",
			"name": "user",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"maxSelect": null,
				"displayFields": null
			}
		}`), del_user)
		collection.Schema.AddField(del_user)

		// add
		del_url := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "lgfvbvbz",
			"name": "url",
			"type": "url",
			"required": true,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), del_url)
		collection.Schema.AddField(del_url)

		// remove
		collection.Schema.RemoveField("xst7gj3i")

		return dao.SaveCollection(collection)
	})
}
