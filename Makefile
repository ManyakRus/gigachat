SERVICENAME=gigachat
SERVICEURL=github.com/ManyakRus/$(SERVICENAME)

FILEMAIN=./cmd/gigachat/main.go
FILEAPP=./bin/gigachat

NEW_REPO=$(SERVICEURL)


run:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	#	cd ./bin && \
	$(FILEAPP)
mod:
	clear
	go mod tidy -compat=1.18
	go mod vendor
	go fmt ./...
build:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	cd ./cmd && \
	./VersionToFile.py
buildwin:
	cls
	go build -race -o bin\gigachat.exe internal\v0\app\main.go
	cd cmd && \
	VersionToFile.py
lint:
	clear
	go fmt ./...
	golangci-lint run ./internal/...
	golangci-lint run ./pkg/...
	gocyclo -over 10 ./internal
	gocyclo -over 10 ./pkg
	gocritic check ./internal/...
	gocritic check ./pkg/...
	staticcheck ./internal/...
	staticcheck ./pkg/...
run.test:
	clear
	go fmt ./...
	go test -coverprofile cover.out ./internal/... ./cmd/...
	go tool cover -func=cover.out
newrepo:
	sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g' go.mod
	find -name *.go -not -path "*/vendor/*"|xargs sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g'
graph:
	clear
	image_packages ./ docs/packages.graphml
conn:
	clear
	image_connections ./internal/v0/app docs/connections.graphml $(SERVICENAME)
