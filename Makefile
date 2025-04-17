BINARY_NAME=deepeye
BINARY_VERSION=$(shell cat version.txt)
BIN_FOLDER=bin
SRC_FOLDER=src
BINARY_FULLNAME=${BINARY_NAME}-${BINARY_VERSION}

# Architecture targets
LINUX_AMD64=${BINARY_FULLNAME}-linux-amd64
LINUX_ARM64=${BINARY_FULLNAME}-linux-arm64
MAC_AMD64=${BINARY_FULLNAME}-darwin-amd64
MAC_ARM64=${BINARY_FULLNAME}-darwin-arm64
WIN_AMD64=${BINARY_FULLNAME}-windows-amd64.exe

# Build all platforms
build-all: clean build-linux build-mac build-windows

# Build for current platform (default)
build: clean
	@mkdir -p ./${BIN_FOLDER}
	@go build -o ./${BIN_FOLDER}/${BINARY_FULLNAME} ./${SRC_FOLDER}/*.go
	@tar -C ./${BIN_FOLDER} -czf ./${BIN_FOLDER}/${BINARY_FULLNAME}.tar.gz ${BINARY_FULLNAME}
	@chmod +x ./${BIN_FOLDER}/${BINARY_FULLNAME}
	@echo "Built ${BINARY_FULLNAME} for current platform"

# Build for Linux (amd64 and arm64)
build-linux: clean
	@mkdir -p ./${BIN_FOLDER}
	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build -o ./${BIN_FOLDER}/${LINUX_AMD64} ./${SRC_FOLDER}/*.go
	@chmod +x ./${BIN_FOLDER}/${LINUX_AMD64}
	@tar -C ./${BIN_FOLDER} -czf ./${BIN_FOLDER}/${LINUX_AMD64}.tar.gz ${LINUX_AMD64}
	
	@echo "Building for Linux (arm64)..."
	@GOOS=linux GOARCH=arm64 go build -o ./${BIN_FOLDER}/${LINUX_ARM64} ./${SRC_FOLDER}/*.go
	@chmod +x ./${BIN_FOLDER}/${LINUX_ARM64}
	@tar -C ./${BIN_FOLDER} -czf ./${BIN_FOLDER}/${LINUX_ARM64}.tar.gz ${LINUX_ARM64}
	
	@echo "Linux builds completed"

# Build for macOS (amd64 and arm64)
build-mac: clean
	@mkdir -p ./${BIN_FOLDER}
	@echo "Building for macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 go build -o ./${BIN_FOLDER}/${MAC_AMD64} ./${SRC_FOLDER}/*.go
	@chmod +x ./${BIN_FOLDER}/${MAC_AMD64}
	@tar -C ./${BIN_FOLDER} -czf ./${BIN_FOLDER}/${MAC_AMD64}.tar.gz ${MAC_AMD64}
	
	@echo "Building for macOS (arm64)..."
	@GOOS=darwin GOARCH=arm64 go build -o ./${BIN_FOLDER}/${MAC_ARM64} ./${SRC_FOLDER}/*.go
	@chmod +x ./${BIN_FOLDER}/${MAC_ARM64}
	@tar -C ./${BIN_FOLDER} -czf ./${BIN_FOLDER}/${MAC_ARM64}.tar.gz ${MAC_ARM64}
	
	@echo "macOS builds completed"

# Build for Windows (amd64)
build-windows: clean
	@mkdir -p ./${BIN_FOLDER}
	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build -o ./${BIN_FOLDER}/${WIN_AMD64} ./${SRC_FOLDER}/*.go
	@zip -j ./${BIN_FOLDER}/${BINARY_FULLNAME}-windows-amd64.zip ./${BIN_FOLDER}/${WIN_AMD64}
	@echo "Windows build completed"

# Create release packages for all platforms
release: build-all
	@mkdir -p ./${BIN_FOLDER}/release
	@cp ./${BIN_FOLDER}/*.tar.gz ./${BIN_FOLDER}/release/
	@cp ./${BIN_FOLDER}/*.zip ./${BIN_FOLDER}/release/
	@cp ./install_linux.sh ./${BIN_FOLDER}/release/
	@cp ./install_mac.sh ./${BIN_FOLDER}/release/
	@cp ./install_win.ps1 ./${BIN_FOLDER}/release/
	@echo "Release packages created in ./${BIN_FOLDER}/release"

version-bump:
	@read -p "Current version: $(shell cat version.txt). Enter new version (without v prefix, e.g. 0.0.3): " newver; \
	echo $$newver > version.txt; \
	echo "Version updated to $$newver"

check:
	@./${BIN_FOLDER}/${BINARY_FULLNAME} -v
	@./${BIN_FOLDER}/${BINARY_FULLNAME} -h

run: build
	./${BIN_FOLDER}/${BINARY_FULLNAME}

install: build
	@sudo rm -rf /usr/local/bin/${BINARY_NAME} && sudo cp ./${BIN_FOLDER}/${BINARY_FULLNAME} /usr/local/bin/${BINARY_NAME}

install-linux: build-linux
	@sudo rm -rf /usr/local/bin/${BINARY_NAME} && sudo cp ./${BIN_FOLDER}/${LINUX_AMD64} /usr/local/bin/${BINARY_NAME}

install-mac: build-mac
	@sudo rm -rf /usr/local/bin/${BINARY_NAME} && sudo cp ./${BIN_FOLDER}/${MAC_AMD64} /usr/local/bin/${BINARY_NAME}

test:
	go test ./src/...

coverage:
	go test -coverpkg=./... ./...

clean:
	@go clean
	@rm -rf ./${BIN_FOLDER}
