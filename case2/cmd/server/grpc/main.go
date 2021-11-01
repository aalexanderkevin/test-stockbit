package main

import (
	"case2/app/database"
	"case2/config"
	handler "case2/handlers"
	"case2/moviepb"
	"case2/repository"
	service "case2/services"
	"case2/thirdparty"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
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

	lis, err := net.Listen("tcp", conf.Grpc.Host+":"+conf.Grpc.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	moviepb.RegisterOmdbServer(server, handler.NewGrpcHandler(s))

	term := make(chan os.Signal)
	go func() {
		log.Println("grpc server listening at " + lis.Addr().String())
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			term <- syscall.SIGINT
		}
	}()

	signal.Notify(term, syscall.SIGTERM, syscall.SIGINT)
	<-term
	log.Println("shutting down")
	server.GracefulStop()
}
