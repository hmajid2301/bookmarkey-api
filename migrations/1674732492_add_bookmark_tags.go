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
			"id": "w4u7mhgxrb2z1nq",
			"created": "2023-01-21 18:42:07.323Z",
			"updated": "2023-01-21 18:42:07.323Z",
			"name": "bookmark_tags",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "zqi0wtxh",
					"name": "bookmark",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"maxSelect": null,
						"collectionId": "tg5uixgnz32nmqj",
						"cascadeDelete": false
					}
				},
				{
					"system": false,
					"id": "g5iqxmew",
					"name": "tag",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"maxSelect": null,
						"collectionId": "0rkrqvkbrwzy4z1",
						"cascadeDelete": false
					}
				}
			],
			"listRule": "@request.auth.id != bookmark.user.id",
			"viewRule": "@request.auth.id != bookmark.user.id",
			"createRule": "@request.auth.id != ''",
			"updateRule": "@request.auth.id != bookmark.user.id",
			"deleteRule": "@request.auth.id != bookmark.user.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("w4u7mhgxrb2z1nq")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
