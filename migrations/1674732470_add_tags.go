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
			"id": "0rkrqvkbrwzy4z1",
			"created": "2023-01-21 18:42:07.322Z",
			"updated": "2023-01-21 18:42:07.322Z",
			"name": "tags",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "3khhvbpc",
					"name": "tag",
					"type": "text",
					"required": true,
					"unique": true,
					"options": {
						"min": 1,
						"max": 100,
						"pattern": ""
					}
				}
			],
			"listRule": "@request.auth.id != ''",
			"viewRule": "@request.auth.id != ''",
			"createRule": "@request.auth.id != ''",
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

		collection, err := dao.FindCollectionByNameOrId("0rkrqvkbrwzy4z1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
