
all: gateway

gateway:
	cd server/cmd/gateway && go build -v

# Run tests
test: fmt vet
	go test ./... -coverprofile cover.out

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

clean:
	rm -rf server/cmd/gateway/gateway
	rm -rf cover.out
