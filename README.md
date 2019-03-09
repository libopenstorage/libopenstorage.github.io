# OpenStorage API Website
This stores the sources and output for https://libopenstorage.github.io. This website is created using
[gitbook](https://github.com/GitbookIO/gitbook-cli). All sources and markdown files are in the `docs` directory. The
output html created by gitbook is located in `w` directory.

## Environment setup
First, you may need to install [npm](https://nodejs.org/en/download/). Then install gitbook:

```
$ npm install -g gitbook-cli
```

To build:

```
make
```

To serve locally on `http://localhost:4000` while editing, run:

```
make serve
```

## Generating API Documentation
To generate it, you will need to install [protoc-gen-doc](https://github.com/lpabon/protoc-gen-doc/releases/tag/v1.1.1)


Then type `make api` to generate the API documentation.

## Releasing
When releasing new version use the following method to maintain that release
version on the documentation website:

1. Add an entry to `getcontent.sh`
1. Edit `reference.md.tmpl`
1. Add an entry in the reference in `docs/SUMMARY.md`
1. Create `docs/swagger-ui/<branch>.index.html` and adjust the json file accordingly

## More information
* [Gitbook](https://toolchain.gitbook.com/)
* Examples gitbook json setup
    * https://github.com/poppy-project/poppy-docs/blob/master/book.json
