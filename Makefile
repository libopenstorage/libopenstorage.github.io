HAS_PROTOC := $(shell command -v protoc 2> /dev/null)
HAS_PROTOGENDOC := $(shell command -v protoc-gen-doc 2> /dev/null)
all: build

images:
	$(MAKE) -C docs/images

build: api changelog images
	gitbook build docs/ w

serve: images
	gitbook serve docs/ w

changelog:
	curl https://raw.githubusercontent.com/libopenstorage/openstorage/master/SDK_CHANGELOG.md \
		--output docs/changelog.md --silent

api:
ifndef HAS_PROTOC
	$(error "Please install protoc 1.3.3 or later")
endif
ifndef HAS_PROTOGENDOC
	$(error "Please install protoc-gen-doc. See README.md for more information")
endif
	curl https://raw.githubusercontent.com/libopenstorage/openstorage/master/api/server/sdk/api/api.swagger.json \
		--output docs/api/api.swagger.json --silent
	curl https://raw.githubusercontent.com/libopenstorage/openstorage/master/api/api.proto \
		--output api.proto --silent
	protoc -I. -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--doc_out=docs/ --doc_opt=./template/sdk.tmpl,generated-api.md api.proto
	rm -f api.proto

# If you have updated any plugins in docs/book.json from https://plugins.gitbook.com/
# you will need to update the modules
update-modules:
	cd docs && rm -rf node_modules && gitbook install


.PHONY: all build serve images api
