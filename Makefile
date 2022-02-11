PREFIX?=/usr/local
GOCMD=go
BINARY_NAME=nextfetch
DESTDIR:=

build:
	$(GOCMD) mod vendor
	GO111MODULE=on $(GOCMD) build -mod=vendor -ldflags="-w -s" -tags=osusergo -o ./out/$(BINARY_NAME)

clean:
	rm -rf ./out
	rm -rf ./vendor

install:
	install -Dm755 ./out/$(BINARY_NAME) $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)

install-config:
	install ./nextfetch.default $(HOME)/.nextfetch

uninstall:
	rm -f $(PREFIX)/bin/$(BINARY_NAME)