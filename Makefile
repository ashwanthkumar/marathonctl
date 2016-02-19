APPNAME = marathonctl
VERSION=0.0.3-fix

build:
	go build -o ${APPNAME} .

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -X main.APP_VERSION=${VERSION}" -v -o ${APPNAME}-linux-amd64 .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -X main.APP_VERSION=${VERSION}" -v -o ${APPNAME}-darwin-amd64 .

build-all: build-mac build-linux

all: setup
	build
	install

setup:  
	go get github.com/spf13/cobra
	go get github.com/spf13/viper
	go get github.com/parnurzeal/gorequest
	go get github.com/gosuri/uiprogress
	# Test deps
	go get github.com/stretchr/testify/assert

test:
	go test github.com/ashwanthkumar/marathonctl/client

install: build
	sudo install -d /usr/local/bin
	sudo install -c ${APPNAME} /usr/local/bin/${APPNAME}

uninstall:
	sudo rm /usr/local/bin/${APPNAME}
