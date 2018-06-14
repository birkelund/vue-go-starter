.PHONY: all clean dev
.PHONY: apiserver ui 

DEP := $(shell command -v dep 2> /dev/null)
STATIK := $(shell command -v statik 2> /dev/null)


all: apiserver

dev:
	go run main.go -http=:4000

ui:
	$(MAKE) -C ui

statik: ui
ifndef STATIK
	$(error "please install statik; go get github.com/rakyll/statik")
endif
	go generate ./...

apiserver: statik
ifndef DEP
	$(error "please install dep; go get -u github.com/golang/dep/cmd/dep")
endif
	dep ensure
	go generate ./...
	go build

clean:
	rm -f vue-go-starter
	$(MAKE) -C ui clean

cleanall: clean
	rm -rf statik
	rm -rf vendor
	$(MAKE) -C ui cleanall
