//go:build integration
// +build integration

package collections

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
	"github.com/pocketbase/pocketbase/tokens"
)

const testDataDir = "../../tests/pb_data"

func TestDeleteCollection(t *testing.T) {
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
			Name:   "Successfully delete collection",
			Method: http.MethodDelete,
			Url:    "/collections/c8qow1f695xqg84",
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"name", "software", "message", "Successfully deleted collection."},
			ExpectedEvents:  map[string]int{"OnModelAfterDelete": 1, "OnModelBeforeDelete": 1},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "Fail to delete collection not authenticated",
			Method:          http.MethodDelete,
			Url:             "/collections/gtx1lh0l990bxnw",
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"message", "The request requires valid record authorization token to be set."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Fail to delete collection belonging to another user",
			Method: http.MethodDelete,
			Url:    "/collections/gtx1lh0l990bxnw",
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"message", "The user does not have permission to delete collection"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestAddCollection(t *testing.T) {
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
			Name:   "Successfully add collection",
			Method: http.MethodPost,
			Url:    "/collections",
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			Body:            strings.NewReader(`{"collection_name": "newCollection"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"name", "newCollection", "message", "Successfully created collection."},
			ExpectedEvents:  map[string]int{"OnModelAfterCreate": 1, "OnModelBeforeCreate": 1},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Fail to add collection incorrect field name body",
			Method: http.MethodPost,
			Url:    "/collections",
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			Body:            strings.NewReader(`{"collection": "newCollection"}`),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"message", "Missing `collection_name` field in request."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Fail to add collection no body",
			Method: http.MethodPost,
			Url:    "/collections",
			RequestHeaders: map[string]string{
				"Authorization": recordToken,
			},
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"message", "Failed to decode payload when trying to add collection."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "Fail to add collection not authenticated",
			Method:          http.MethodPost,
			Url:             "/collections",
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"message", "The request requires valid record authorization token to be set."},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

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
