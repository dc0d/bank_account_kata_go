.PHONY: test
test:
	clear
	go test -count=1 -timeout 60s -p 1 -coverprofile=./cover/all-profile.out -covermode=set -coverpkg=./... ./...; \
	go tool cover -html=./cover/all-profile.out -o ./cover/all-coverage.html

lint:
	golangci-lint run ./...

gen:
	go generate ./...
