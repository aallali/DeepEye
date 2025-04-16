#!/bin/bash
set -e  # Exit immediately if a command exits with non-zero status

# Get version from remote version file
echo "Fetching latest version information..."
VERSION_URL="https://raw.githubusercontent.com/aallali/DeepEye/main/version.txt"
VERSION=$(wget -qO- ${VERSION_URL} || curl -s ${VERSION_URL})

if [ -z "$VERSION" ]; then
    echo "ERROR: Failed to retrieve version information"
    exit 1
fi

EXECUTABLE_VERSION="${VERSION}"
EXECUTABLE_NAME="deepeye-${EXECUTABLE_VERSION}"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="deepeye"
DOWNLOAD_URL="https://github.com/aallali/DeepEye/releases/download/${EXECUTABLE_VERSION}/${EXECUTABLE_NAME}-mac.tar.gz"

echo "Latest version: ${EXECUTABLE_VERSION}"
echo "Download URL: ${DOWNLOAD_URL}"

# Check if running as root
if [ "$(id -u)" -ne 0 ]; then
    echo "ERROR: This script must be run as root"
    exit 1
fi

echo "Downloading DeepEye ${EXECUTABLE_VERSION}..."
if command -v wget > /dev/null; then
    wget -q -O "${EXECUTABLE_NAME}.tar.gz" "${DOWNLOAD_URL}" || 
        { echo "Failed to download DeepEye"; exit 1; }
elif command -v curl > /dev/null; then
    curl -s -L -o "${EXECUTABLE_NAME}.tar.gz" "${DOWNLOAD_URL}" || 
        { echo "Failed to download DeepEye"; exit 1; }
else
    echo "ERROR: Neither wget nor curl is installed. Please install one of them and try again."
    exit 1
fi

echo "Extracting..."
tar -xf "${EXECUTABLE_NAME}.tar.gz" || 
    { echo "Failed to extract DeepEye"; exit 1; }

echo "Installing to ${INSTALL_DIR}/${BINARY_NAME}..."
# Remove old binary if exists
rm -f "${INSTALL_DIR}/${BINARY_NAME}"
# Install new binary
install -m 755 "${EXECUTABLE_NAME}-linux-amd64" "${INSTALL_DIR}/${BINARY_NAME}" || 
    { echo "Failed to install DeepEye"; exit 1; }

# Clean up
rm -f "${EXECUTABLE_NAME}.tar.gz" "${EXECUTABLE_NAME}-linux-amd64"

echo "DeepEye ${EXECUTABLE_VERSION} installed successfully!"
echo "Run 'deepeye -h' for usage information"

cat << EOF
To enable automatic update checks, add to your shell config:

For Bash users:
  echo "deepeye -u" >> ~/.bashrc

For Zsh users:
  echo "deepeye -u" >> ~/.zshrc

EOF