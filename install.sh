#!/bin/bash

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

echo "DeepEye Installer"
echo "Detecting system: $OS ($ARCH)"

case "$OS" in
    linux*)
        echo "✓ Detected Linux system"
        if [ "$(id -u)" -ne 0 ]; then
            echo "This installation requires root privileges. Executing with sudo..."
            sudo bash install.sh
        else
            bash install.sh
        fi
        ;;
    darwin*)
        echo "✓ Detected macOS system"
        if [ "$(id -u)" -ne 0 ]; then
            echo "This installation requires root privileges. Executing with sudo..."
            sudo bash install_mac.sh
        else
            bash install_mac.sh
        fi
        ;;
    msys*|mingw*|cygwin*)
        echo "✓ Detected Windows system"
        echo "Windows installation requires PowerShell with administrator privileges."
        echo ""
        echo "Please run the following in an Administrator PowerShell window:"
        echo "    .\\install.ps1"
        echo ""
        echo "Tip: Right-click PowerShell and select 'Run as Administrator'"
        ;;
    *)
        echo "❌ Unsupported operating system: $OS"
        echo "Please manually install DeepEye for your platform"
        exit 1
        ;;
esac