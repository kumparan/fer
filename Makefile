SHELL:=/bin/bash
changelog_args=-o CHANGELOG.md -p '^v'

proto:
	@protoc --go_out=plugins=grpc:. pb/example/*.proto
	@ls pb/example/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

create-version:
ifndef version
	$(error version is not set)
endif
	@echo -e 'package config\n\n// Version define version of fer' >config/version.go
	@echo 'const Version = "$(version)"' >>config/version.go

run:
	@go run main.go

build:
	@go build -o ./bin/fer

changelog: fetch-git create-changelog create-version

fetch-git:
	@echo 'fetching remote'
	@git fetch

create-changelog:
ifndef version
	$(error version is not set)
endif
	$(eval changelog_args=--next-tag $(version) $(changelog_args))
	git-chglog $(changelog_args)

lint:
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --enable=golint --enable=goimports  --enable=unconvert --enable=unparam --concurrency=2


