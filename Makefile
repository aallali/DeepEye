BINARY_NAME=deepeye
BINARY_VERSION=v0.0.2
BIN_FOLDER=bin
SRC_FOLDER=src
BINARY_FULLNAME=${BINARY_NAME}-${BINARY_VERSION}

build: clean
	@rm -rf ./${BIN_FOLDER}/*
	@go build -o ./${BIN_FOLDER}/${BINARY_FULLNAME} ./${SRC_FOLDER}/*.go
 
	@tar -C ./${BIN_FOLDER}  -czf ./${BIN_FOLDER}/${BINARY_FULLNAME}.tar.gz ${BINARY_FULLNAME}
	@chmod +x ./${BIN_FOLDER}/${BINARY_FULLNAME}
 

check:
	@./${BIN_FOLDER}/${BINARY_FULLNAME} -v
	@./${BIN_FOLDER}/${BINARY_FULLNAME} -h

run: build
	./${BIN_FOLDER}/${BINARY_FULLNAME}

install: build
	@sudo rm -rf /usr/local/bin/${BINARY_NAME} && sudo cp ./${BIN_FOLDER}/${BINARY_FULLNAME} /usr/local/bin/${BINARY_NAME}

test:
	go test ./src/...

coverage:
	go test -coverpkg=./... ./...

clean:
	@go clean
	@rm -rf ./${BIN_FOLDER}
 