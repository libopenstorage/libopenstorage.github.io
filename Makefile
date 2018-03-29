all: build

build:
	gitbook build docs/ w

serve:
	gitbook serve docs/ w


.PHONY: all build serve
