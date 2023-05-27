BINARY_NAME=deepeye
BIN_FOLDER=bin
SRC_FOLDER=src

build:
	@go build -o ./${BIN_FOLDER}/${BINARY_NAME} ./${SRC_FOLDER}/*.go

run: build
	./${BIN_FOLDER}/${BINARY_NAME}

runc: 
	@clear && make run

install: build
	@rm -rf /usr/local/bin/${BINARY_NAME} && chmod 777 ./${BIN_FOLDER}/${BINARY_NAME} && sudo cp ./${BIN_FOLDER}/${BINARY_NAME} /usr/local/bin/

clean:
	@go clean
	@rm -rf ./${BIN_FOLDER}
 