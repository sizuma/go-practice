BINARY_NAME=../../bin/c2c
CONFIG_SAMPLE_NAME=./resource/config.sample.toml
CONFIG_NAME=./resource/config.toml
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
	go get github.com/go-sql-driver/mysql
	go get github.com/BurntSushi/toml
	go get github.com/Masterminds/squirrel
	-@cp -n $(CONFIG_SAMPLE_NAME) $(CONFIG_NAME)

build: main.go
	make clean
	make check
	go build -o $(BINARY_NAME)
run: $(BINARY_NAME) $(CONFIG_NAME)
	make build
	$(BINARY_NAME)
test-run:
	go build -o $(BINARY_NAME)
	$(BINARY_NAME)