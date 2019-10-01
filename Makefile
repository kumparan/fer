proto:
	@protoc --go_out=plugins=grpc:. pb/content/*.proto
	@ls pb/content/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

run:
	@go run main.go

build:
	@go build -o ./bin/fer

changelog:
	@git-chglog -o CHANGELOG.md

lint:
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --enable=golint --enable=goimports  --enable=unconvert --enable=unparam --concurrency=2


