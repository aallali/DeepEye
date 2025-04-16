#!/bin/bash

set -e  # Exit immediately if a command exits with non-zero status

# Get version from remote version file
echo "Fetching latest version information..."
VERSION_URL="https://raw.githubusercontent.com/aallali/DeepEye/main/version.txt"
VERSION=$(curl -s ${VERSION_URL})

if [ -z "$VERSION" ]; then
    echo "ERROR: Failed to retrieve version information"
    exit 1
fi

EXECUTABLE_VERSION="0.0.3"
EXECUTABLE_NAME="deepeye-${EXECUTABLE_VERSION}"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="deepeye"
DOWNLOAD_URL="https://github.com/aallali/DeepEye/releases/download/${EXECUTABLE_VERSION}/${EXECUTABLE_NAME}-darwin-amd64.tar.gz"

echo "Latest version: ${EXECUTABLE_VERSION}"
echo "Download URL: ${DOWNLOAD_URL}"

# Check if running as root or with sudo
if [ "$(id -u)" -ne 0 ]; then
    echo "ERROR: This script must be run with sudo"
    exit 1
fi

echo "Downloading DeepEye ${EXECUTABLE_VERSION}..."
curl -s -L -o "${EXECUTABLE_NAME}.tar.gz" "${DOWNLOAD_URL}" || 
    { echo "Failed to download DeepEye"; exit 1; }

echo "Extracting..."
tar -xf "${EXECUTABLE_NAME}.tar.gz" || 
    { echo "Failed to extract DeepEye"; exit 1; }

echo "Installing to ${INSTALL_DIR}/${BINARY_NAME}..."
# Remove old binary if exists
rm -f "${INSTALL_DIR}/${BINARY_NAME}"
# Install new binary
install -m 755 "${EXECUTABLE_NAME}-darwin-amd64" "${INSTALL_DIR}/${BINARY_NAME}" || 
    { echo "Failed to install DeepEye"; exit 1; }

# Clean up
rm -f "${EXECUTABLE_NAME}.tar.gz" "${EXECUTABLE_NAME}-darwin-amd64"

echo "DeepEye ${EXECUTABLE_VERSION} installed successfully!"
echo "Run 'deepeye -h' for usage information"

cat << EOF
To enable automatic update checks, add to your shell config:

For Bash users:
  echo "deepeye -u" >> ~/.bash_profile

For Zsh users:
  echo "deepeye -u" >> ~/.zshrc

EOF