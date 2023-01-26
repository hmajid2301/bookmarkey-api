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
			"id": "tg5uixgnz32nmqj",
			"created": "2023-01-21 18:42:07.322Z",
			"updated": "2023-01-25 23:05:50.970Z",
			"name": "bookmarks",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "_pb_users_auth_",
					"name": "user",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"maxSelect": null,
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false
					}
				},
				{
					"system": false,
					"id": "sabinmy4",
					"name": "collection",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"maxSelect": null,
						"collectionId": "34588opuh85l19p",
						"cascadeDelete": false
					}
				},
				{
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
				},
				{
					"system": false,
					"id": "kmbgnjco",
					"name": "favourite",
					"type": "bool",
					"required": true,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "k2lfxned",
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
			"listRule": "@request.auth.id = user.id",
			"viewRule": "@request.auth.id = user.id",
			"createRule": "@request.auth.id != ''",
			"updateRule": "@request.auth.id = user.id",
			"deleteRule": "@request.auth.id = user.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
