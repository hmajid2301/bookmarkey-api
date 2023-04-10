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

		collection, err := dao.FindCollectionByNameOrId("xowwnq4hswfdsci")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_xowwnq4hswfdsci_created_idx` + "`" + ` ON ` + "`" + `groups` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_yR9lUni` + "`" + ` ON ` + "`" + `groups` + "`" + ` (` + "`" + `user` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("xowwnq4hswfdsci")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_xowwnq4hswfdsci_created_idx` + "`" + ` ON ` + "`" + `groups` + "`" + ` (` + "`" + `created` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	})
}
