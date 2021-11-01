package service

import (
	model "case2/models"
	"case2/repository"
	"case2/thirdparty"
	"context"
	"encoding/json"
	"time"
)

type Movie interface {
	MovieSearch(ctx context.Context, req model.SearchRequest) (*model.SearchResponse, error)
	MovieDetail(ctx context.Context, req model.DetailRequest) (*model.Movie, error)
}

type movie struct {
	Repository repository.Repository
	ThirdParty thirdparty.MovieThirdParty
}

func NewMovieService(r repository.Repository, tp thirdparty.MovieThirdParty) Movie {
	return &movie{
		Repository: r,
		ThirdParty: tp,
	}
}

func (m *movie) MovieSearch(ctx context.Context, req model.SearchRequest) (*model.SearchResponse, error) {
	logger := &model.Logger{
		Timestamp: time.Now(),
		Request:   req.Request,
	}
	var err error
	res, err := m.ThirdParty.Search(ctx, req.Search, req.Page)
	jsonRes, e := json.Marshal(res)
	if e != nil {
		logger.Response = "Error"
	}

	logger.Response = string(jsonRes)
	go m.Repository.Insert(*logger)

	return res, err
}

func (m *movie) MovieDetail(ctx context.Context, req model.DetailRequest) (*model.Movie, error) {
	m.Repository.Insert(model.Logger{
		Timestamp: time.Now(),
	})

	res, err := m.ThirdParty.GetDetail(ctx, req.Id)
	if err != nil {
		return res, err
	}

	return res, nil
}
