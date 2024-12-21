#!/usr/bin/env bash

set -e

echo "Installing Hard CLI..."

# Check if curl is installed
if ! command -v curl &> /dev/null; then
    echo "Error: curl is not installed. Please install curl and try again."
    exit 1
fi

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Error: Docker is not installed. Please install Docker and try again."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
    echo "Error: Docker Compose is not installed. Please install Docker Compose and try again."
    exit 1
fi

# Download the latest hard binary
echo "Downloading the latest Hard CLI binary..."
LATEST_URL=$(curl -s https://api.github.com/repos/clebsonsh/hard/releases/latest | grep "browser_download_url.*hard\"" | cut -d : -f 2,3 | tr -d \")
curl -L -o hard "$LATEST_URL"

# Make the binary executable
chmod +x hard

# Move the binary to a directory in PATH
echo "Moving Hard CLI to /usr/local/bin (may require sudo)..."
sudo mv hard /usr/local/bin/

echo "Hard CLI installation complete!"
echo "You can now use the 'hard' command to manage your PHP development environment."
echo "For more information, run 'hard --help'."
