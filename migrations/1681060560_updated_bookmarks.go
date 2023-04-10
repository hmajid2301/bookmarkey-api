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

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_tg5uixgnz32nmqj_created_idx` + "`" + ` ON ` + "`" + `bookmarks` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_xXZ0Flu` + "`" + ` ON ` + "`" + `bookmarks` + "`" + ` (` + "`" + `user` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tg5uixgnz32nmqj")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_tg5uixgnz32nmqj_created_idx` + "`" + ` ON ` + "`" + `bookmarks` + "`" + ` (` + "`" + `created` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	})
}
