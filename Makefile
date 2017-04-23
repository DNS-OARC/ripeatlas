SOURCES := $(wildcard *.go) \
    $(wildcard measurement/*.go) \
    $(wildcard measurement/dns/*.go)

EXAMPLE_SOURCES = $(wildcard example/*.go)

all: .built

fmt: format

clean:
	rm -f .built

format:
	gofmt -w $(SOURCES) $(EXAMPLE_SOURCES)
	sed -i -e 's%	%    %g' $(SOURCES) $(EXAMPLE_SOURCES)

.built: $(SOURCES)
	go build -v -x
	touch .built
