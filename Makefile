test: test_01 test_02 test_04

test_01:
	@cd ./internal/convertor && go test . -v

test_02:
	@cd ./internal/analys && go test . -v

test_04:
	@cd ./internal/monitors && go test . -v

run:
	@go run ./cmd/app/