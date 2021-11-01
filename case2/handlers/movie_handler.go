package handler

import (
	model "case2/models"
	"case2/moviepb"
	service "case2/services"
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type httpHandler struct {
	s service.Movie
}

type grpcHandler struct {
	moviepb.UnimplementedOmdbServer

	s service.Movie
}

func NewHandler(movie service.Movie) *httpHandler {
	return &httpHandler{
		s: movie,
	}
}

func NewGrpcHandler(movie service.Movie) *grpcHandler {
	return &grpcHandler{
		s: movie,
	}
}

func (h httpHandler) HandleSearchMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.SearchRequest{
			Request: r.URL.String(),
			Page:    r.URL.Query().Get("page"),
			Search:  r.URL.Query().Get("search"),
		}
		movie, err := h.s.MovieSearch(r.Context(), req)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}
		WriteSuccessResponse(w, http.StatusOK, movie, nil)
	}
}

func (h httpHandler) HandleGetMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		req := model.DetailRequest{
			Id: params["id"],
		}
		movie, err := h.s.MovieDetail(r.Context(), req)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}
		WriteSuccessResponse(w, http.StatusOK, movie, nil)
	}
}

func (h grpcHandler) MovieSearch(ctx context.Context, r *moviepb.SearchRequest) (*moviepb.SearchResponse, error) {
	req := model.SearchRequest{
		Page:   r.GetPage(),
		Search: r.GetSearch(),
	}

	movie, err := h.s.MovieSearch(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	response := &moviepb.SearchResponse{
		Total:        movie.Total,
		ErrorMessage: movie.ErrorMessage,
	}

	for _, val := range movie.Search {
		response.Search = append(response.Search, &moviepb.Movie{
			Title:   val.Title,
			Year:    val.Year,
			MovieId: val.MovieID,
			Type:    val.Type,
			Poster:  val.Poster,
		})
	}

	return response, nil
}

func (h grpcHandler) MovieDetail(ctx context.Context, r *moviepb.DetailRequest) (*moviepb.Movie, error) {
	req := model.DetailRequest{
		Id: r.GetId(),
	}

	movie, err := h.s.MovieDetail(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	response := &moviepb.Movie{
		Title:   movie.Title,
		Year:    movie.Year,
		MovieId: movie.MovieID,
		Type:    movie.Type,
		Poster:  movie.Poster,
	}

	return response, nil
}
