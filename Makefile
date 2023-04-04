name = $(shell basename $(CURDIR))
target = $(name)
sources = $(name).go
archbuild = \
	mkdir -p build/$(1)_$(2); \
	GOARCH=$(2) GOOS=$(1) go build -o build/$(1)_$(2)/$(name)$(if $(filter windows,$(1)),.exe) $(sources); \
	cp README.md build/$(1)_$(2); \
	$(if $(filter tar,$(3)), \
	tar cvzf release/$(name)_$(version)_$(1)_$(2).tar.gz -C build $(1)_$(2) \
	, \
	cd build; zip -r ../release/$(name)_$(version)_$(1)_$(2).zip $(1)_$(2) )
version = $(shell git describe --always --tags --dirty)

.PHONY: test test_coverage clean release install

all: $(target)

test:
	go test ./...

test_coverage:
	go test -cover ./...

clean:
	rm -f $(target)

$(target): $(sources)
	go build -o $@ $^

release: test $(sources) 
	mkdir -p release
	$(call archbuild,darwin,arm64,tar)
	$(call archbuild,darwin,amd64,tar)
	$(call archbuild,linux,arm64,tar)
	$(call archbuild,linux,amd64,tar)
	$(call archbuild,windows,arm64,zip)
	$(call archbuild,windows,amd64,zip)

install: $(target)
	go install -v .
