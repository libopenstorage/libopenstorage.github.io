# OpenStorage API Website
This stores the sources and output for https://libopenstorage.github.io.

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
1. Add an entry in the reference in `mkdocs.yml'
1. Create `docs/swagger-ui/<branch>.index.html` and adjust the json file accordingly

