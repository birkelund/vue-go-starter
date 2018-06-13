.PHONY: build-deps deps all clean build

YARN := $(shell command -v yarn 2> /dev/null)


all: build-deps deps build

build-deps:
ifndef YARN
	$(error "please install yarn; https://yarnpkg.com")
endif

deps:
	yarn
build:
	yarn build

clean:
	rm -rf dist

cleanall: clean
	rm -rf node_modules
