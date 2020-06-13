# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test build
build:
	rm -rf target/
	mkdir target/
	cp cmd/*.toml target/
	$(GOBUILD) -o target/web cmd/main.go

test:
	$(GOTEST) -v ./...

clean:
	rm -rf target/

run:
	target/web -env=test -conf=target/web_conf.toml 2>&1 > target/web.log

stop:
	pkill -f target/web

doc:
	swag init -g cmd/main.go

run-back:
	nohup target/web -env=test -conf=target/web_conf.toml 2>&1 > target/web.log &

xorm-model:
	xorm reverse mysql "root:e23456@tcp(192.168.1.162:3307)/test?charset=utf8" dbmodel/goxorm dbmodel
