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

		collection, err := dao.FindCollectionByNameOrId("grh3maryyyl7zm1")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_grh3maryyyl7zm1_created_idx` + "`" + ` ON ` + "`" + `nested_collections` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_mmnNXsM` + "`" + ` ON ` + "`" + `nested_collections` + "`" + ` (` + "`" + `parent_collection` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("grh3maryyyl7zm1")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_grh3maryyyl7zm1_created_idx` + "`" + ` ON ` + "`" + `nested_collections` + "`" + ` (` + "`" + `created` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	})
}
