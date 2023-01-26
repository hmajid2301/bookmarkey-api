package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "xowwnq4hswfdsci",
			"created": "2023-01-21 19:11:06.955Z",
			"updated": "2023-01-25 23:06:29.172Z",
			"name": "groups",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "gqkol4wb",
					"name": "user",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false
					}
				},
				{
					"system": false,
					"id": "ryl5rl8w",
					"name": "name",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": 1,
						"max": 256,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "fkc0oxbe",
					"name": "custom_order",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null
					}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("xowwnq4hswfdsci")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
