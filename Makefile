.PHONY: build format lint doc validate-openapi

build:
	go build

format:
	go fmt ./...

lint:

doc:

validate-openapi:
	openapi-generator validate --input-spec ./openapi.yml --recommend
