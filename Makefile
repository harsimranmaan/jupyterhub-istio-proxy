test: lint
	go test -coverprofile=coverage.txt -covermode=atomic -race ./...
lint: vet
	golangci-lint run
vet:
	go vet ./...