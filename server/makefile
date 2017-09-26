GOCMD   := go
GOBUILD := $(GOCMD) build
GOTEST  := $(GOCMD) test

all: test build

build:
	$(GOBUILD)
	sudo setcap CAP_NET_BIND_SERVICE=+ep server

test:
	$(GOTEST)

certs:
	openssl req -new -nodes -x509 -newkey rsa -keyout ssl/private/key.pem -out ssl/certs/cert.pem -subj "/"
