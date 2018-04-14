BINARY_NAME=../../bin/c2c

.PHONY: test

default:
	make build
clean:
	@rm -f "$(BINARY_NAME)"
test:
	go test c2c/test/...
vet:
	go vet c2c/...
lint:
	golint -set_exit_status=1 c2c/...
fmt:
	find ./ -name "*.go" | xargs gofmt -w
check:
	make vet
	make lint
	make fmt
	make test
prepare:
build: main.go
	make clean
	make check
	go build -o $(BINARY_NAME)
run: $(BINARY_NAME)
	make build
	$(BINARY_NAME)