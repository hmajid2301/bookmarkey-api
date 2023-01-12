// +gocover:ignore:file ignore this file!
package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	isUserLoggedIn := "@request.auth.id != ''"
	isUser := "@request.auth.id != user"
	collections := []*models.Collection{
		{
			Name:       "tags",
			Type:       models.CollectionTypeBase,
			ListRule:   types.Pointer(isUserLoggedIn),
			ViewRule:   types.Pointer(isUserLoggedIn),
			CreateRule: types.Pointer(isUserLoggedIn),
			UpdateRule: nil,
			DeleteRule: nil,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "tag",
					Type:     schema.FieldTypeText,
					Required: true,
					Unique:   true,
					Options: &schema.TextOptions{
						Min: types.Pointer(1),
						Max: types.Pointer(100),
					},
				},
			),
		},

		{
			Name:       "bookmarks",
			Type:       models.CollectionTypeBase,
			ListRule:   types.Pointer(isUser),
			ViewRule:   types.Pointer(isUser),
			CreateRule: types.Pointer(isUserLoggedIn),
			UpdateRule: types.Pointer(isUser),
			DeleteRule: types.Pointer(isUser),
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "user",
					Type:     schema.FieldTypeRelation,
					Required: true,
					Options: &schema.RelationOptions{
						MaxSelect:    nil,
						CollectionId: "users",
					},
				},
				&schema.SchemaField{
					Name:     "collection",
					Type:     schema.FieldTypeRelation,
					Required: true,
					Options: &schema.RelationOptions{
						MaxSelect:    nil,
						CollectionId: "collections",
					},
				},
				&schema.SchemaField{
					Name:     "url",
					Type:     schema.FieldTypeUrl,
					Required: true,
					Options:  &schema.UrlOptions{},
				},
				&schema.SchemaField{
					Name:     "favourite",
					Type:     schema.FieldTypeBool,
					Required: true,
					Options:  &schema.BoolOptions{},
				},
			),
		},

		{
			Name:       "collections",
			Type:       models.CollectionTypeBase,
			ListRule:   types.Pointer(isUser),
			ViewRule:   types.Pointer(isUser),
			CreateRule: types.Pointer(isUserLoggedIn),
			UpdateRule: types.Pointer(isUser),
			DeleteRule: types.Pointer(isUser),
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "parent",
					Type:     schema.FieldTypeRelation,
					Required: false,
					Options: &schema.RelationOptions{
						MaxSelect:    nil,
						CollectionId: "collections",
					},
				},
				&schema.SchemaField{
					Name:     "name",
					Type:     schema.FieldTypeText,
					Required: true,
					Options: &schema.TextOptions{
						Min: types.Pointer(1),
						Max: types.Pointer(256),
					},
				},
				&schema.SchemaField{
					Name:     "user",
					Type:     schema.FieldTypeRelation,
					Required: true,
					Options: &schema.RelationOptions{
						MaxSelect:    nil,
						CollectionId: "users",
					},
				},
			),
		},

		{
			Name:       "bookmark_tags",
			Type:       models.CollectionTypeBase,
			ListRule:   types.Pointer("@request.auth.id != bookmark.user.id"),
			ViewRule:   types.Pointer("@request.auth.id != bookmark.user.id"),
			CreateRule: types.Pointer(isUserLoggedIn),
			UpdateRule: types.Pointer("@request.auth.id != bookmark.user.id"),
			DeleteRule: types.Pointer("@request.auth.id != bookmark.user.id"),
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "bookmark",
					Type:     schema.FieldTypeRelation,
					Required: true,
					Options: &schema.RelationOptions{
						MaxSelect:    nil,
						CollectionId: "bookmarks",
					},
				},
				&schema.SchemaField{
					Name:     "tag",
					Type:     schema.FieldTypeRelation,
					Required: true,
					Options: &schema.RelationOptions{
						MaxSelect:    nil,
						CollectionId: "tags",
					},
				},
			),
		},
	}

	m.Register(func(db dbx.Builder) error {
		for i := 0; i < len(collections); i++ {
			err := daos.New(db).SaveCollection(collections[i])
			if err != nil {
				return err
			}
		}
		return nil

	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		for i := 0; i < len(collections); i++ {
			collection, err := dao.FindCollectionByNameOrId(collections[i].Name)
			if err != nil {
				return err
			}

			err = daos.New(db).DeleteCollection(collection)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
