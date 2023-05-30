BINARY_NAME=deepeye
BINARY_VERSION=0.0.1
BIN_FOLDER=bin
SRC_FOLDER=src
BINARY_FULLNAME=${BINARY_NAME}-${BINARY_VERSION}

build: clean
	@rm -rf ./${BIN_FOLDER}/*
	@go build -o ./${BIN_FOLDER}/${BINARY_FULLNAME} ./${SRC_FOLDER}/*.go
	@zip ./${BIN_FOLDER}/${BINARY_FULLNAME}.zip ./${BIN_FOLDER}/${BINARY_FULLNAME} 
	@chmod 777 ./${BIN_FOLDER}/${BINARY_FULLNAME}

check:
	@./${BIN_FOLDER}/${BINARY_FULLNAME} -v
	@./${BIN_FOLDER}/${BINARY_FULLNAME} -h

run: build
	./${BIN_FOLDER}/${BINARY_FULLNAME}

install: build
	@sudo rm -rf /usr/local/bin/${BINARY_NAME} && sudo cp ./${BIN_FOLDER}/${BINARY_FULLNAME} /usr/local/bin/${BINARY_NAME}

clean:
	@go clean
	@rm -rf ./${BIN_FOLDER}
 