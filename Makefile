build:
	go build -o bin/mini-k8s

run: build
	./bin/mini-k8s

test:
	go test -v ./...