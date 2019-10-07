SHELL:=/bin/bash
changelog_args=-o CHANGELOG.md -p '^v'

proto:
	@protoc --go_out=plugins=grpc:. pb/example/*.proto
	@ls pb/example/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

run:
	@go run main.go

build:
	@go build -o ./bin/fer

changelog:
ifdef version
	$(eval changelog_args=--next-tag $(version) $(changelog_args))
endif
	git-chglog $(changelog_args)


lint:
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --enable=golint --enable=goimports  --enable=unconvert --enable=unparam --concurrency=2


