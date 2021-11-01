runhttp:
	go run ./cmd/server/http/main.go
	
rungrpc:
	go run ./cmd/server/grpc/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./moviepb/movie.proto