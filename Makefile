default: test

test:
	go test ./...

release:
	goreleaser --rm-dist

ci: test
