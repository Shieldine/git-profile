# Define variables
$repo = "Shieldine/git-profile"
$installDir = "$Env:UserProfile\AppData\Local\Programs\git-profile"
$exeName = "git-profile.exe" # Name of the executable

# Create installation directory if it doesn't exist
if (-Not (Test-Path -Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir
}

# Determine system architecture
$arch = if ([System.Environment]::Is64BitOperatingSystem) {
    "x86_64"
} else {
    "i386"
}

# Check if the system is ARM64
if ([System.Environment]::ProcessorArchitecture -eq "ARM64") {
    $arch = "arm64"
}

Write-Output "Architecture $arch detected"

# get the latest release information from GitHub API
$release = (Invoke-RestMethod -Uri "https://api.github.com/repos/$repo/releases/latest").tag_name
if (-not $release) {
    Write-Host "Error: unable to retrieve latest release tag"
    exit 1
}

# Download the asset
$zipUrl = "https://github.com/$repo/releases/download/$release/git-profile_Windows_$arch.zip"
$zipFile = "$installDir\git-profile_Windows_$arch.zip"
Invoke-WebRequest -Uri $zipUrl -OutFile $zipFile

# Extract the zip file
Add-Type -AssemblyName System.IO.Compression.FileSystem
[System.IO.Compression.ZipFile]::ExtractToDirectory($zipFile, $installDir)

# Remove the zip file after extraction
Remove-Item -Path $zipFile

# Add the installation directory to the system PATH
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$installDir", [System.EnvironmentVariableTarget]::User)
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","User")

# Verify that the executable is added to the PATH
if (Get-Command $exeName -ErrorAction SilentlyContinue) {
    Write-Output "Installation successful. $exeName is now in the PATH."
} else {
    Write-Output "Installation successful, but adding $exeName to PATH has failed. Please set it manually."
}
