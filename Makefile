all: build

build:
	gitbook build docs/ w

serve:
	gitbook serve docs/ w

# If you have updated any plugins in docs/book.json from https://plugins.gitbook.com/
# you will need to update the modules
update-modules:
	cd docs && rm -rf node_modules && gitbook install


.PHONY: all build serve
