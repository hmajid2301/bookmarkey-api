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
			"id": "pejwlxcamufi2z9",
			"created": "2023-01-21 18:42:07.322Z",
			"updated": "2023-01-25 23:06:12.162Z",
			"name": "collections",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "awpaxjtv",
					"name": "parent",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"collectionId": "pejwlxcamufi2z9",
						"cascadeDelete": false
					}
				},
				{
					"system": false,
					"id": "v7sahbbo",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 1,
						"max": 256,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "y7fety8o",
					"name": "user",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false
					}
				},
				{
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
				},
				{
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

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
