GOCMD   := go
GOBUILD := $(GOCMD) build
GOTEST  := $(GOCMD) test

all: test certs build

build:
	$(GOBUILD)
	sudo setcap cap_net_bind_service=+ep server

test:
	$(GOTEST)

certs:
	openssl req -new -nodes -x509 -newkey rsa -keyout server.key -out server.crt -subj "/"

clean:
	$(RM) server*
