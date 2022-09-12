HAS_PROTOC := $(shell command -v protoc 2> /dev/null)
HAS_PROTOGENDOC := $(shell command -v protoc-gen-doc 2> /dev/null)
all: build

images:
	$(MAKE) -C docs/images

build: api images docs/reference.md
	gitbook build docs/ w

serve: api images docs/reference.md
	gitbook serve docs/ w

api:
ifndef DOCKER_PROTO
	docker run \
        --privileged --rm \
        -v $(shell pwd):/go/src/code \
        -e "GOPATH=/go" \
        -e "DOCKER_PROTO=yes" \
        -e "PROTO_USER=$(shell id -u)" \
        -e "PROTO_GROUP=$(shell id -g)" \
        -e "PATH=/bin:/usr/bin:/usr/local/bin:/go/bin:/usr/local/go/bin" \
        quay.io/openstorage/grpc-framework:latest \
            make api
else
	bash getcontent.sh
endif

docs/reference.md: api

# If you have updated any plugins in docs/book.json from https://plugins.gitbook.com/
# you will need to update the modules
update-modules:
	cd docs && rm -rf node_modules && gitbook install


.PHONY: all build serve images api
