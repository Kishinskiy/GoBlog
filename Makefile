.PHONY: build
build:
	go build -o ./bin/myapp cmd/goservice/main.go

.PHONY: run
run: build
	./bin/myapp