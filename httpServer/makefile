APP=http_server
.PHONY: build
build: clean
	go build -o ${APP} http_server.go
.PHONY: run
run:
	go run -race http_server.go
.PHONY: clean
clean:
	go clean
.PHONY: api
api:
	protoc --go_out=../protocal/ --go-grpc_out=../protocal/ --go-grpc_opt=require_unimplemented_servers=false --proto_path=../protocal/  ../protocal/et_tcp.proto