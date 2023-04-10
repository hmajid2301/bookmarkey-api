package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_pejwlxcamufi2z9_created_idx` + "`" + ` ON ` + "`" + `collections` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_CreVYUz` + "`" + ` ON ` + "`" + `collections` + "`" + ` (` + "`" + `user` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_hT589R8` + "`" + ` ON ` + "`" + `collections` + "`" + ` (` + "`" + `group` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pejwlxcamufi2z9")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_pejwlxcamufi2z9_created_idx` + "`" + ` ON ` + "`" + `collections` + "`" + ` (` + "`" + `created` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	})
}
