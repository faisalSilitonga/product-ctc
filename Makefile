.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin
BINARY=bin/engine

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies ; go build -tags $(LIBRARY_ENV) -o ${BINARY} cmd/main.go

engine:
	go build -tags $(LIBRARY_ENV) -o ${BINARY} cmd/main.go

run: build ; ./${BINARY}

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/app api/main.go

ci: dependencies test	

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/product/interface.go -destination=usecase/product/mock/product.go -package=mock

test:
	go test -tags testing ./...

# docker:
# 	docker build -t product-ctc .

# docker-run:
# 	docker-compose up --build -d

# docker-stop:
# 	docker-compose down

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done