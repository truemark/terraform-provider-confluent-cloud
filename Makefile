.PHONY: build clean install go-format go-lint go-doc tf-doc

build:
	./bin/build.sh

clean:
	./bin/clean.sh

install:
	./bin/install.sh


go-format:
	go fmt ./...

go-lint:


go-doc:


tf-doc:

