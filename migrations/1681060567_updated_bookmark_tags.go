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

		collection, err := dao.FindCollectionByNameOrId("w4u7mhgxrb2z1nq")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_w4u7mhgxrb2z1nq_created_idx` + "`" + ` ON ` + "`" + `bookmark_tags` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_tuLF8zZ` + "`" + ` ON ` + "`" + `bookmark_tags` + "`" + ` (` + "`" + `bookmark` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("w4u7mhgxrb2z1nq")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_w4u7mhgxrb2z1nq_created_idx` + "`" + ` ON ` + "`" + `bookmark_tags` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_tuLF8zZ` + "`" + ` ON ` + "`" + `bookmark_tags` + "`" + ` (` + "`" + `bookmark` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	})
}
