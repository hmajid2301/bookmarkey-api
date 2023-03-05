//go:build integration
// +build integration

package bookmarks

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
	"github.com/pocketbase/pocketbase/tokens"
)

const testDataDir = "../../tests/pb_data"

func TestCreateBookmark(t *testing.T) {
	recordToken, err := generateRecordToken("users", "test@bookmarkey.app")
	if err != nil {
		t.Fatal(err)
	}

	setupTestApp := func() (*tests.TestApp, error) {
		testApp, err := tests.NewTestApp(testDataDir)
		if err != nil {
			return nil, err
		}

		AddHandlers(testApp)
		return testApp, nil
	}

	scenarios := []tests.ApiScenario{
		{
			Name:   "Successfully create bookmark",
			Method: http.MethodPost,
			Url:    "/collections/6oj42javvobz2fx/bookmark",
			Body: strings.NewReader(`{
				"url": "https://blog.pragmaticengineer.com/"
			}`),
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			ExpectedStatus: http.StatusCreated,
			ExpectedEvents: map[string]int{"OnModelAfterCreate": 2, "OnModelBeforeCreate": 2},
			TestAppFactory: setupTestApp,
		},
		{
			Name:   "Successfully create bookmark, update bookmark metadata url exists",
			Method: http.MethodPost,
			Url:    "/collections/6oj42javvobz2fx/bookmark",
			Body: strings.NewReader(`{
				"url": "https://wiki.guildwars2.com/wiki/Event_timers"
			}`),
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			ExpectedStatus: http.StatusCreated,
			ExpectedEvents: map[string]int{"OnModelAfterCreate": 1, "OnModelAfterUpdate": 1, "OnModelBeforeCreate": 1, "OnModelBeforeUpdate": 1},
			TestAppFactory: setupTestApp,
		},
		{
			Name:   "Fail to create bookmark invalid payload",
			Method: http.MethodPost,
			Url:    "/collections/6oj42javvobz2fx/bookmark",
			Body: strings.NewReader(`{
				"wrong_field": "https://blog.pragmaticengineer.com/"
			}`),
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			ExpectedContent: []string{"message", "Key: 'NewBookmark.URL' Error:Field validation for 'URL' failed on the 'required' tag."},
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedEvents:  map[string]int{},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Fail to create bookmark unauthorized",
			Method: http.MethodPost,
			Url:    "/collections/6oj42javvobz2fx/bookmark",
			Body: strings.NewReader(`{
				"url": "https://blog.pragmaticengineer.com/"
			}`),
			RequestHeaders:  map[string]string{},
			ExpectedContent: []string{"message", "The request requires valid record authorization token to be set."},
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedEvents:  map[string]int{},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

// TODO: generalise
func generateRecordToken(collectionNameOrId string, email string) (string, error) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		return "", err
	}
	defer app.Cleanup()

	record, err := app.Dao().FindAuthRecordByEmail(collectionNameOrId, email)
	if err != nil {
		return "", err
	}

	return tokens.NewRecordAuthToken(app, record)
}
