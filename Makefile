build:
	go build .

run:
	go run main.go

tests:
	go test ./...

build-docker:
	docker build -t ricardo:v0.1 .

run-docker:
	docker run --rm -it ricardo:v0.1

run:
	docker-compose up -d