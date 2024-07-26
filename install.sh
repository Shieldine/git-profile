#!/bin/bash

REPO="Shieldine/git-profile"
INSTALL_DIR="$HOME/bin"

OS=$(uname -s)
ARCH=$(uname -m)

case "$OS" in
    "Darwin")
        case "$ARCH" in
            "x86_64")
                BINARY_SUFFIX="Darwin_x86_64"
                ;;
            "arm64")
                BINARY_SUFFIX="Darwin_arm64"
                ;;
            *)
                echo "Unsupported architecture: $ARCH on macOS"
                exit 1
                ;;
        esac
        ;;
    "Linux")
        case "$ARCH" in
            "x86_64")
                BINARY_SUFFIX="Linux_x86_64"
                ;;
            "arm64" | "aarch64")
                BINARY_SUFFIX="Linux_arm64"
                ;;
            "i386")
                BINARY_SUFFIX="Linux_i386"
                ;;
            *)
                echo "Unsupported architecture: $ARCH on Linux"
                exit 1
                ;;
        esac
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Fetch the latest release tag from GitHub
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | awk -F '"tag_name": "' '{if ($2) print $2}' | awk -F'"' '{print $1}')

if [ -z "$LATEST_RELEASE" ]; then
    echo "Failed to fetch the latest release"
    exit 1
fi

# Construct the archive name and URL
ARCHIVE="git-profile_$BINARY_SUFFIX.tar.gz"
URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/$ARCHIVE"

# Download the archive
echo "Downloading $ARCHIVE from $URL..."
curl -L -o /tmp/"$ARCHIVE" "$URL"

# Extract the archive
echo "Extracting $ARCHIVE..."
mkdir -p /tmp/git-profile
tar -xzf /tmp/"$ARCHIVE" -C /tmp/git-profile

# move to install dir
mv /tmp/git-profile/git-profile "$INSTALL_DIR"/git-profile
echo "Executable moved to $INSTALL_DIR"

# clean up tmp files
rm /tmp/"$ARCHIVE"
rm -rf /tmp/git-profile/

# make git-profile executable
chmod +x "$INSTALL_DIR"/git-profile

# Add install dir to PATH if not already there
[[ ":$PATH:" != *":$INSTALL_DIR:"* ]] && PATH="$INSTALL_DIR:${PATH}"

# Check if the Fish shell is in use and add to Fish's user paths
if [ "$FISH_VERSION" = "" ]; then
  echo "Fish shell detected. Adding to fish path..."
  fish -c "fish_add_path $INSTALL_DIR"
fi

# Verify installation
if command -v git-profile &> /dev/null; then
    echo "Installation successful!"
else
    echo "Installation failed."
fi
