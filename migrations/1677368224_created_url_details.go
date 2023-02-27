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
			"id": "elul616p7fivid4",
			"created": "2023-02-25 23:37:04.845Z",
			"updated": "2023-02-25 23:37:04.845Z",
			"name": "url_details",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jpw7marq",
					"name": "url",
					"type": "url",
					"required": true,
					"unique": false,
					"options": {
						"exceptDomains": [],
						"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "1efxsron",
					"name": "image",
					"type": "url",
					"required": true,
					"unique": false,
					"options": {
						"exceptDomains": [],
						"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "5kccyw2t",
					"name": "description",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 0,
						"max": 250,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "hpzhvex4",
					"name": "title",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 0,
						"max": 100,
						"pattern": ""
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

		collection, err := dao.FindCollectionByNameOrId("elul616p7fivid4")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
