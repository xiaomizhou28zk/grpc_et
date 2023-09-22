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