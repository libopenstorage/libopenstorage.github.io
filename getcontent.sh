#!/bin/sh

BRANCHES="master release-9.2 release-7.0 release-6.4 release-4.0 release-sdk-0.9"

getBranch() {
	local branch=$1

	echo ">>> Collecting information from branch ${branch}"
	curl https://raw.githubusercontent.com/libopenstorage/openstorage/${branch}/api/server/sdk/api/api.swagger.json \
		--output docs/api/${branch}.api.swagger.json --silent
	curl https://raw.githubusercontent.com/libopenstorage/openstorage/${branch}/SDK_CHANGELOG.md \
		--output docs/${branch}.changelog.md --silent
	curl https://raw.githubusercontent.com/libopenstorage/openstorage/${branch}/api/api.proto \
		--output ${branch}.api.proto --silent
	protoc -I. -I /usr/local/include -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--doc_out=./template/sdk.tmpl,${branch}.generated-api.md:docs/ ${branch}.api.proto
	rm -f ${branch}.api.proto

	ver=$(cat docs/api/${branch}.api.swagger.json | jq -r '.info.version')

	VERSIONS="$VERSIONS -e s#@@${branch}-version@@#$ver#g"
}

for b in ${BRANCHES} ; do
	getBranch $b
done

echo ">>> Created docs/reference.md"
sed $VERSIONS reference.md.tmpl > docs/reference.md
