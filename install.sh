#!/bin/bash

REPO="Shieldine/git-profile"
INSTALL_DIR="$HOME/.local/bin/git-profile"

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

if [ $? -ne 0 ]; then
    echo "error: failed to download $ARCHIVE from $URL"
    exit 1
fi

echo "Extracting $ARCHIVE to $INSTALL_DIR..."

# create install dir
mkdir -p "$INSTALL_DIR"
if [ $? -ne 0 ]; then
    echo "error: unable to create install dir $INSTALL_DIR"
    exit 1
fi

# extract archive
tar -xzf /tmp/"$ARCHIVE" -C "$INSTALL_DIR"
if [ $? -ne 0 ]; then
    echo "error: unable to extract $ARCHIVE"
    exit 1
fi

# clean up tmp file
rm /tmp/"$ARCHIVE"

# make git-profile executable
chmod +x "$INSTALL_DIR"/git-profile
if [ $? -ne 0 ]; then
    echo "error: unable to make binary executable"
    exit 1
fi

read -r -p "Do you want to add the program to PATH? (y/n): " user_response
user_response=$(echo "$user_response" | tr '[:upper:]' '[:lower:]')


if [[ "$user_response" == "y" || "$user_response" == "yes" ]]; then

  [[ ":$PATH:" != *":$INSTALL_DIR:"* ]] && PATH="$INSTALL_DIR:${PATH}"
  echo "Added $INSTALL_DIR to PATH."

  if command -v fish &> /dev/null; then
    echo "Fish shell detected. Adding to fish path..."
    fish -c "fish_add_path $INSTALL_DIR"
  fi

else
  echo "The program was not added to PATH."
fi

echo "Installation finished."
