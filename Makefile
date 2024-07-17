build-gin:
	go build ./cmd/gin/gin.go
build-http:
	go build ./cmd/http/http.go

format:
	gofmt ./

lint:
	golangci-lint run ./...

test:
	go test -cover ./pkg/...

build-docker-images:
	docker-compose build

run-docker-images:
	docker-compose up

