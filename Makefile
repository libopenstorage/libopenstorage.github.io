all: build

images:
	$(MAKE) -C docs/images

build: images
	gitbook build docs/ w

serve: images
	gitbook serve docs/ w

# If you have updated any plugins in docs/book.json from https://plugins.gitbook.com/
# you will need to update the modules
update-modules:
	cd docs && rm -rf node_modules && gitbook install


.PHONY: all build serve images
