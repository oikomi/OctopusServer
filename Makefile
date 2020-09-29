
ifneq "$(strip $(shell command -v go 2>/dev/null))" ""
	GOOS ?= $(shell go env GOOS)
	GOARCH ?= $(shell go env GOARCH)
else
	ifeq ($(GOOS),)
		# approximate GOOS for the platform if we don't have Go and GOOS isn't
		# set. We leave GOARCH unset, so that may need to be fixed.
		ifeq ($(OS),Windows_NT)
			GOOS = windows
		else
			UNAME_S := $(shell uname -s)
			ifeq ($(UNAME_S),Linux)
				GOOS = linux
			endif
			ifeq ($(UNAME_S),Darwin)
				GOOS = darwin
			endif
			ifeq ($(UNAME_S),FreeBSD)
				GOOS = freebsd
			endif
		endif
	else
		GOOS ?= $$GOOS
		GOARCH ?= $$GOARCH
	endif
endif


#Replaces ":" (*nix), ";" (windows) with newline for easy parsing
GOPATHS=$(shell echo ${GOPATH} | tr ":" "\n" | tr ";" "\n")

DESTDIR ?= ./deploy

GO_BUILD_FLAGS=

all: clean conf gateway

conf:
	@mkdir -p $(DESTDIR)/config
	@mkdir -p $(DESTDIR)/bin
	@cp config/*.yaml $(DESTDIR)/config

gateway:
	@go build ${GO_BUILD_FLAGS} -o $(DESTDIR)/bin/gateway  ./server/cmd/gateway

#install: ## install binaries
#	@echo "$(WHALE) $@ $(BINARIES)"
#	@mkdir -p $(DESTDIR)/bin
#	@install $(BINARIES) $(DESTDIR)/bin
#
#uninstall: clean
#	@echo "$(WHALE) $@"
#	@rm -f $(addprefix $(DESTDIR)/bin/,$(notdir $(BINARIES)))

# Run tests
test: fmt vet
	@go test ./... -coverprofile cover.out

# Run go fmt against code
fmt:
	@go fmt ./...

# Run go vet against code
vet:
	@go vet ./...

clean:
	@rm -rf deploy
	@rm -rf cover.out
