// Package helpers provides a bunch of common/helper functions such as simplifying testing (reducing boilerplate)
package helpers

import (
	"github.com/pocketbase/pocketbase/tests"
	"github.com/pocketbase/pocketbase/tokens"
)

// TestDataDir relative path to folder containing test database file
const TestDataDir = "../../tests/pb_data"

// GenerateRecordToken returns a JWT token we can use for testing
func GenerateRecordToken(email string) (string, error) {
	app, err := tests.NewTestApp(TestDataDir)
	if err != nil {
		return "", err
	}
	defer app.Cleanup()

	record, err := app.Dao().FindAuthRecordByEmail("users", email)
	if err != nil {
		return "", err
	}

	return tokens.NewRecordAuthToken(app, record)
}
