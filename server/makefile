GOCMD   := go
GOBUILD := $(GOCMD) build

build:
	$(GOBUILD) -o build/server
	sudo setcap CAP_NET_BIND_SERVICE=+ep build/server
	openssl req -new -nodes -x509 -newkey rsa -keyout build/key.pem -out build/cert.pem -subj "/"
