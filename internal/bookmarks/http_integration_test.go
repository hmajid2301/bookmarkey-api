//go:build integration
// +build integration

package bookmarks

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
	"gitlab.com/bookmarkey/api/internal/helpers"
	"gitlab.com/bookmarkey/api/internal/middleware"
)

func TestCreateBookmark(t *testing.T) {
	jwtToken := getJWTToken(t, "test@bookmarkey.app")
	setupTestApp := setupTests(t)

	scenarios := []tests.ApiScenario{
		{
			Name:   "Successfully create bookmark",
			Method: http.MethodPost,
			Url:    "/collections/6oj42javvobz2fx/bookmark",
			Body: strings.NewReader(`{
				"url": "https://blog.pragmaticengineer.com/"
			}`),
			RequestHeaders: map[string]string{
				"Authorization": jwtToken,
			},
			ExpectedStatus: http.StatusCreated,
			ExpectedEvents: map[string]int{"OnModelAfterCreate": 2, "OnModelBeforeCreate": 2},
			TestAppFactory: setupTestApp,
		},
		{
			Name:   "Successfully create bookmark, unsorted",
			Method: http.MethodPost,
			Url:    "/collections/-1/bookmark",
			Body: strings.NewReader(`{
				"url": "https://blog.pragmaticengineer.com/"
			}`),
			RequestHeaders: map[string]string{
				"Authorization": jwtToken,
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
				"Authorization": jwtToken,
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
				"Authorization": jwtToken,
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

func TestGetURLMetadata(t *testing.T) {
	jwtToken := getJWTToken(t, "test@bookmarkey.app")
	setupTestApp := setupTests(t)

	scenarios := []tests.ApiScenario{
		{
			Name:   "Successfully get URL metadata",
			Method: http.MethodGet,
			Url:    "/bookmark/metadata?url=" + url.QueryEscape("https://blog.pragmaticengineer.com/"),
			RequestHeaders: map[string]string{
				"Authorization": jwtToken,
			},
			ExpectedContent: []string{"url", "https://blog.pragmaticengineer.com/", "description", "Observations across the software engineering industry.", "title", "The Pragmatic Engineer", "image", ""},
			ExpectedStatus:  http.StatusOK,
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Successfully get URL metadata, that exists in DB",
			Method: http.MethodGet,
			Url:    "/bookmark/metadata?url=" + url.QueryEscape("https://stackoverflow.com/questions/7172784/how-do-i-post-json-data-with-curl"),
			RequestHeaders: map[string]string{
				"Authorization": jwtToken,
			},
			ExpectedContent: []string{"url", "https://stackoverflow.com/questions/7172784/how-do-i-post-json-data-with-curl", "description", "I use Ubuntu and installed cURL on it. I want to test my Spring REST application with cURL. I wrote my POST code at the Java side. However, I want to test it with cURL. I am trying to post a JSON d...", "title", "How do I POST JSON data with cURL?", "image", "https://cdn.sstatic.net/Sites/stackoverflow/Img/apple-touch-icon@2.png?v=73d79a89bded"},
			ExpectedStatus:  http.StatusOK,
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Fail to get URL metadata, invalid URL",
			Method: http.MethodGet,
			Url:    "/bookmark/metadata?url=" + url.QueryEscape("https://"),
			RequestHeaders: map[string]string{
				"Authorization": jwtToken,
			},
			ExpectedContent: []string{"message", "Expected valid URL in query parameter."},
			ExpectedStatus:  http.StatusBadRequest,
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Fail to get URL metadata, missing URL",
			Method: http.MethodGet,
			Url:    "/bookmark/metadata",
			RequestHeaders: map[string]string{
				"Authorization": jwtToken,
			},
			ExpectedContent: []string{"message", "Expected valid URL in query parameter."},
			ExpectedStatus:  http.StatusBadRequest,
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Fail to get URL metadata not authorized",
			Method: http.MethodGet,
			Url:    "/bookmark/metadata?url=" + url.QueryEscape("https://blog.pragmaticengineer.com/"),
			RequestHeaders: map[string]string{
				"Authorization": "",
			},
			ExpectedContent: []string{"message", "The request requires valid record authorization token to be set."},
			ExpectedStatus:  http.StatusUnauthorized,
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func getJWTToken(t *testing.T, userEmail string) string {
	jwtToken, err := helpers.GenerateRecordToken(userEmail)
	if err != nil {
		t.Fatal(err)
	}

	return jwtToken

}

func setupTests(t *testing.T) func() (*tests.TestApp, error) {
	setupTestApp := func() (*tests.TestApp, error) {
		testApp, err := tests.NewTestApp(helpers.TestDataDir)
		middleware.ApplyMiddleware(testApp.BaseApp)
		if err != nil {
			return nil, err
		}

		AddHandlers(testApp)
		return testApp, nil
	}

	return setupTestApp
}
