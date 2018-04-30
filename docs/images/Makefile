HAS_DITAA := $(shell command -v ditaa 2> /dev/null)

DITAA_IMAGES := arch.png sdk.png
DITAA_SRCS := $(patsubst %.png,%.ditaa,$(DITAA_IMAGES))

all: $(DITAA_IMAGES)

%.png: %.ditaa
ifndef HAS_DITAA
	$(error "Please install ditaa to create images")
endif
	$(info $(DITAA_SRCS))
	ditaa -o $<

clean:
	rm -f $(DITAA_IMAGES)

.PHONY: all
