.PHONY: test
test:
	go test -v ./...

.PHONY: test-integration
test-integration:
	cd tests/integration; go test -v -count=1 ./...
