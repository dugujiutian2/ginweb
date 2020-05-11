# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test build
build:
	rm -rf target/
	mkdir target/
	cp cmd/web_conf.json target/web_conf.json
	$(GOBUILD) -o target/web cmd/main.go

test:
	$(GOTEST) -v ./...

clean:
	rm -rf target/

run:
	nohup target/web -env=test -conf=target/web_conf.json 2>&1 > target/web.log &

stop:
	pkill -f target/web

