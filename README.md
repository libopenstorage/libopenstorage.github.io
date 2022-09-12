# OpenStorage API Website
This stores the sources and output for https://libopenstorage.github.io. This website is created using
[gitbook](https://github.com/GitbookIO/gitbook-cli). All sources and markdown files are in the `docs` directory. The
output html created by gitbook is located in `w` directory.

## Environment setup

### Install githbook

First, you may need to install [npm](https://nodejs.org/en/download/). Then install gitbook:

```
npm install -g gitbook-cli
```

## Build the website

To build:

```
make
```

To serve locally on `http://localhost:4000` while editing, run:

```
make serve
```

## Releasing
When releasing new version use the following method to maintain that release
version on the documentation website:

1. Add an entry to `getcontent.sh`
1. Edit `reference.md.tmpl`
1. Add an entry in the reference in `docs/SUMMARY.md`
1. Create `docs/swagger-ui/<branch>.index.html` and adjust the json file accordingly

