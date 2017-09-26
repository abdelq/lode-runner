GOCMD   := go
GOBUILD := $(GOCMD) build

build:
	$(GOBUILD)
	sudo setcap CAP_NET_BIND_SERVICE=+ep server

certs:
	openssl req -new -nodes -x509 -newkey rsa -keyout ssl/private/key.pem -out ssl/certs/cert.pem -subj "/"
