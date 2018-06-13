.PHONY: all clean
.PHONY: apiserver ui

DEP := $(shell command -v dep 2> /dev/null)
STATIK := $(shell command -v statik 2> /dev/null)

all: apiserver

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

cleanall:
	rm -rf statik
	$(MAKE) -C ui cleanall
