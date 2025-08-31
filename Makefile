.PHONY: build
build:
	go build -o ./bin/myapp cmd/myapp/main.go

.PHONY: run
run: build
	./bin/myapp