package thirdparty

import (
	"case2/config"
	model "case2/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type MovieThirdParty interface {
	Search(ctx context.Context, search string, page string) (*model.SearchResponse, error)
	GetDetail(ctx context.Context, id string) (*model.Movie, error)
}

type omdb struct {
	baseURL string
	apiKey  string
}

func NewOMDB(conf config.OmdbConfig) MovieThirdParty {
	return &omdb{
		baseURL: "http://www.omdbapi.com",
		apiKey:  conf.ApiKey,
	}
}

func (o *omdb) Search(ctx context.Context, search string, page string) (*model.SearchResponse, error) {
	r := &model.SearchResponse{}
	resp, err := o.requestAPI("search", search, page)
	if err != nil {
		r.ErrorMessage = err.Error()
		return r, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (o *omdb) GetDetail(ctx context.Context, id string) (*model.Movie, error) {
	r := &model.Movie{}
	resp, err := o.requestAPI("id", id)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (o *omdb) requestAPI(apiCategory string, params ...string) (resp []byte, err error) {
	var URL *url.URL
	URL, err = url.Parse(o.baseURL)
	if err != nil {
		return nil, err
	}

	URL.Path += "/"
	parameters := url.Values{}
	parameters.Add("apikey", o.apiKey)

	switch apiCategory {
	case "search":
		parameters.Add("s", params[0])
		parameters.Add("page", params[1])
	case "id":
		parameters.Add("i", params[0])
	}

	URL.RawQuery = parameters.Encode()
	log.Printf("Sending request to %v\n", URL)

	res, err := http.Get(URL.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	log.Printf("Response: [%s] %d", body, res.StatusCode)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed status_code %d received", res.StatusCode)
	}

	return body, nil
}
