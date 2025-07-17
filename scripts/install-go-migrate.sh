#!/usr/bin/env bash

if command -v migrate &> /dev/null; then
  echo "migrate already installed"
  exit 0
fi

set -e

# Detect OS
OS=$(uname -s)
ARCH=$(uname -m)

# Normalize OS
case "$OS" in
  "Linux") OS="linux" ;;
  "Darwin") OS="darwin" ;;
  "MINGW"*|"MSYS"*|"CYGWIN"*|"Windows_NT") OS="windows" ;;
  *) echo "Unsupported OS: $OS" && exit 1 ;;
esac

# Normalize ARCH
case "$ARCH" in
  "x86_64"|"amd64") ARCH="x86_64" ;;
  "arm64"|"aarch64") ARCH="aarch64" ;;
  *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac

VERSION="v4.17.1"
FILE="migrate.${OS}-${ARCH}"
URL="https://github.com/golang-migrate/migrate/releases/download/${VERSION}/${FILE}.tar.gz"

echo "Downloading migrate from $URL ..."
curl -L "$URL" -o migrate.tar.gz

tar -xzf migrate.tar.gz

chmod +x migrate
sudo mv migrate /usr/local/bin/migrate

rm migrate.tar.gz

echo "✅ Installed migrate:"
migrate -version
