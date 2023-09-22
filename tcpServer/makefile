APP=tcp_server
.PHONY: build
build: clean
	go build -o ${APP} tcp_server.go
.PHONY: run
run:
	go run -race tcp_server.go
.PHONY: clean
clean:
	go clean