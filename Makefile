SOURCES := $(wildcard *.go)

all: .built

fmt: format

clean:
	rm -f .built

format:
	gofmt -w *.go example/*.go
	sed -i -e 's%	%    %g' *.go example/*.go

.built: $(SOURCES)
	go build -v -x
	touch .built
