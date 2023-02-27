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
		jsonData := `[
			{
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
			},
			{
				"id": "tg5uixgnz32nmqj",
				"created": "2023-01-21 18:42:07.322Z",
				"updated": "2023-02-26 18:52:21.393Z",
				"name": "bookmarks",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "sabinmy4",
						"name": "collection",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "pejwlxcamufi2z9",
							"cascadeDelete": false,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "kmbgnjco",
						"name": "favourite",
						"type": "bool",
						"required": false,
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
							"max": 9007199254740991
						}
					},
					{
						"system": false,
						"id": "xst7gj3i",
						"name": "bookmark_metadata",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "elul616p7fivid4",
							"cascadeDelete": false,
							"maxSelect": 1,
							"displayFields": []
						}
					}
				],
				"listRule": "@request.auth.id = collection.user.id",
				"viewRule": "@request.auth.id = collection.user.id",
				"createRule": "@request.auth.id != ''",
				"updateRule": "@request.auth.id = collection.user.id",
				"deleteRule": "@request.auth.id = collection.user.id",
				"options": {}
			},
			{
				"id": "pejwlxcamufi2z9",
				"created": "2023-01-21 18:42:07.322Z",
				"updated": "2023-01-29 16:13:04.867Z",
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
							"collectionId": "pejwlxcamufi2z9",
							"cascadeDelete": false,
							"maxSelect": 1,
							"displayFields": null
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
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"maxSelect": 1,
							"displayFields": null
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
							"collectionId": "xowwnq4hswfdsci",
							"cascadeDelete": false,
							"maxSelect": 1,
							"displayFields": null
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
							"max": 9007199254740991
						}
					}
				],
				"listRule": "@request.auth.id = user.id",
				"viewRule": "@request.auth.id = user.id",
				"createRule": "@request.auth.id != ''",
				"updateRule": "@request.auth.id = user.id",
				"deleteRule": "@request.auth.id = user.id",
				"options": {}
			},
			{
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
							"collectionId": "tg5uixgnz32nmqj",
							"cascadeDelete": false,
							"maxSelect": null,
							"displayFields": null
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
							"collectionId": "0rkrqvkbrwzy4z1",
							"cascadeDelete": false,
							"maxSelect": null,
							"displayFields": null
						}
					}
				],
				"listRule": "@request.auth.id != bookmark.user.id",
				"viewRule": "@request.auth.id != bookmark.user.id",
				"createRule": "@request.auth.id != ''",
				"updateRule": "@request.auth.id != bookmark.user.id",
				"deleteRule": "@request.auth.id != bookmark.user.id",
				"options": {}
			},
			{
				"id": "xowwnq4hswfdsci",
				"created": "2023-01-21 19:11:06.955Z",
				"updated": "2023-01-29 16:13:04.867Z",
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
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"maxSelect": 1,
							"displayFields": null
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
							"max": 9007199254740991
						}
					}
				],
				"listRule": "@request.auth.id = user.id",
				"viewRule": "@request.auth.id = user.id",
				"createRule": "@request.auth.id != ''",
				"updateRule": "@request.auth.id = user.id",
				"deleteRule": "@request.auth.id = user.id",
				"options": {}
			},
			{
				"id": "_pb_users_auth_",
				"created": "2023-01-29 16:13:04.861Z",
				"updated": "2023-01-29 16:13:04.861Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "elul616p7fivid4",
				"created": "2023-02-25 23:37:04.845Z",
				"updated": "2023-02-26 18:52:21.396Z",
				"name": "bookmarks_metadata",
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
				"createRule": "@request.auth.id != \"\"",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
