test:
	@go list ./... | while read -r PACKAGE; do go test $${PACKAGE} -cover -v; done

run:
	@go run ./cmd/app/