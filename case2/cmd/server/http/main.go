package main

import (
	"case2/app/database"
	"case2/config"
	handler "case2/handlers"
	"case2/repository"
	service "case2/services"
	"case2/thirdparty"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	conf, err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}

	db := database.GetConnection(conf.DB)
	defer db.Close()

	s := service.NewMovieService(
		repository.NewLogRepository(db),
		thirdparty.NewOMDB(conf.Omdb),
	)
	h := handler.NewHandler(s)

	r := mux.NewRouter()
	r.HandleFunc("/movie", h.HandleSearchMovie()).Methods("GET")
	r.HandleFunc("/movie/{id}", h.HandleGetMovie()).Methods("GET")
	r.Use(loggingMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         conf.Rest.Host + ":" + conf.Rest.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("http server listening at " + srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request: " + r.Method + " " + r.RequestURI + " " + r.Proto)
		next.ServeHTTP(w, r)
	})
}
