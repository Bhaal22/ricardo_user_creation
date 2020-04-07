deps:
	GO111MODULE=on go get github.com/golang/mock/mockgen@latest

build-docker:
	docker build -t ricardo:v0.1 .

run-docker:
	docker run --rm -it ricardo:v0.1