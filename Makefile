

all: gateway

gateway:
	cd server/cmd/gateway && go build -v

clean:
	rm -rf server/cmd/gateway/gateway

