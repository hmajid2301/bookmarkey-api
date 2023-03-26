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
			"id": "grh3maryyyl7zm1",
			"created": "2023-03-26 15:34:29.754Z",
			"updated": "2023-03-26 15:34:29.754Z",
			"name": "nested_collections",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "4603saav",
					"name": "parent_collection",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "pejwlxcamufi2z9",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "6rshnefq",
					"name": "child_collection",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "pejwlxcamufi2z9",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
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
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("grh3maryyyl7zm1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
