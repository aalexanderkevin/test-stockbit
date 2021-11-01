package thirdparty_test

import (
	"context"
	"strings"
	"testing"

	"case2/config"
	. "case2/thirdparty"
)

func TestInvalidApiKey(t *testing.T) {
	omdb := NewOMDB(config.OmdbConfig{
		ApiKey: "",
	})
	_, err := omdb.Search(context.Background(), "batman", "2")
	if err == nil {
		t.Errorf("Expected to fail")
	}
	if err != nil {
		expectedErrorMsg := "failed status_code 401 received"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected- %s, Got- %s", expectedErrorMsg, err)
		}
	}

	_, err = omdb.GetDetail(context.Background(), "tt0137523")
	if err == nil {
		t.Errorf("Expected to fail")
	}
	if err != nil {
		expectedErrorMsg := "failed status_code 401 received"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected- %s, Got- %s", expectedErrorMsg, err)
		}
	}
}

func TestSearch(t *testing.T) {
	omdb := NewOMDB(config.OmdbConfig{
		ApiKey: "faf7e5bb",
	})

	tests := []struct {
		search string
		page   string
	}{
		{
			"Fight Club",
			"1",
		},
		{
			"Batman",
			"2",
		},
		{
			"Superman",
			"2",
		},
	}

	for i, item := range tests {
		resp, err := omdb.Search(context.Background(), item.search, item.page)
		if err != nil {
			t.Errorf("Test[%d]: %s", i, err)
			continue
		}
		if !strings.Contains(resp.Search[0].Title, item.search) {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.search, resp.Search[0].Title)
			continue
		}
	}
}

func TestGetDetail(t *testing.T) {
	omdb := NewOMDB(config.OmdbConfig{
		ApiKey: "faf7e5bb",
	})

	tests := []struct {
		id    string
		title string
		year  string
	}{
		{
			"tt0137523",
			"Fight Club",
			"1999",
		},
		{
			"tt2884018",
			"Macbeth",
			"2015",
		},
	}

	for i, item := range tests {
		resp, err := omdb.GetDetail(context.Background(), item.id)
		if err != nil {
			t.Errorf("Test[%d]: %s", i, err)
			continue
		}
		if resp.Title != item.title {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.title, resp.Title)
			continue
		}
		if resp.Year != item.year {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.year, resp.Year)
			continue
		}
	}
}
