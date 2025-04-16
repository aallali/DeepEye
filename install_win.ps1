$ErrorActionPreference = "Stop"

# Get version from remote file
Write-Host "Fetching latest version information..." -ForegroundColor Cyan
$versionUrl = "https://raw.githubusercontent.com/aallali/DeepEye/main/version.txt"

try {
    [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
    $version = (Invoke-WebRequest -Uri $versionUrl -UseBasicParsing).Content.Trim()
    
    if ([string]::IsNullOrEmpty($version)) {
        throw "Retrieved version is empty"
    }
} catch {
    Write-Host "ERROR: Failed to retrieve version information: $_" -ForegroundColor Red
    exit 1
}

$executableVersion = "$version"
$executableName = "deepeye-${executableVersion}"
$downloadUrl = "https://github.com/aallali/DeepEye/releases/download/${executableVersion}/${executableName}-windows.zip"
$installDir = "$env:LOCALAPPDATA\DeepEye"
$binaryName = "deepeye.exe"

Write-Host "Latest version: $executableVersion" -ForegroundColor Green
Write-Host "Download URL: $downloadUrl" -ForegroundColor Cyan

# Check if running as admin
if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Host "ERROR: This script must be run as Administrator" -ForegroundColor Red
    Write-Host "Right-click PowerShell and select 'Run as Administrator', then run this script again." -ForegroundColor Yellow
    exit 1
}

# Create installation directory if it doesn't exist
if (-not (Test-Path -Path $installDir)) {
    try {
        New-Item -ItemType Directory -Path $installDir -Force | Out-Null
        Write-Host "Created installation directory: $installDir" -ForegroundColor Green
    } catch {
        Write-Host "ERROR: Failed to create installation directory: $_" -ForegroundColor Red
        exit 1
    }
}

# Download DeepEye
Write-Host "Downloading DeepEye ${executableVersion}..." -ForegroundColor Cyan
$zipPath = "$env:TEMP\$executableName.zip"

try {
    Invoke-WebRequest -Uri $downloadUrl -OutFile $zipPath -UseBasicParsing
    
    if (-not (Test-Path -Path $zipPath)) {
        throw "Download completed but file not found at expected location: $zipPath"
    }
} catch {
    Write-Host "ERROR: Failed to download DeepEye: $_" -ForegroundColor Red
    exit 1
}

# Extract the ZIP file
Write-Host "Extracting..." -ForegroundColor Cyan
try {
    Expand-Archive -Path $zipPath -DestinationPath $env:TEMP -Force
    $extractedExe = "$env:TEMP\$executableName.exe"
    
    if (-not (Test-Path -Path $extractedExe)) {
        # Try to find the executable in case the zip structure is different
        $possibleExes = Get-ChildItem -Path $env:TEMP -Filter "*.exe" -Recurse | Where-Object { $_.Name -like "*deepeye*" }
        
        if ($possibleExes.Count -gt 0) {
            $extractedExe = $possibleExes[0].FullName
            Write-Host "Found executable at: $extractedExe" -ForegroundColor Green
        } else {
            throw "Extraction completed but executable not found in the archive"
        }
    }
} catch {
    Write-Host "ERROR: Failed to extract DeepEye: $_" -ForegroundColor Red
    exit 1
}

# Install the binary
Write-Host "Installing to $installDir\$binaryName..." -ForegroundColor Cyan
Copy-Item -Path $extractedExe -Destination "$installDir\$binaryName" -Force

# Add to PATH if not already
$path = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
if ($path -notlike "*$installDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$path;$installDir", [EnvironmentVariableTarget]::User)
    Write-Host "Added DeepEye to PATH" -ForegroundColor Green
}

# Clean up temporary files
Remove-Item -Path $zipPath -Force -ErrorAction SilentlyContinue
Remove-Item -Path $extractedExe -Force -ErrorAction SilentlyContinue

Write-Host "DeepEye ${executableVersion} installed successfully!" -ForegroundColor Green
Write-Host "Run 'deepeye -h' for usage information" -ForegroundColor Cyan
Write-Host @"

To enable automatic update checks:
    1. Create a shortcut to DeepEye in your Startup folder
    2. Set the target to: $installDir\$binaryName -u

"@ -ForegroundColor Yellow