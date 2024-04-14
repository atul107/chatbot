#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Define colors for pretty output
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Function to display messages
echo_green() {
    echo -e "${GREEN}$1${NC}"
}

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go."
    exit 1
fi

echo_green "Go is installed."

# Install Go dependencies
echo_green "Installing Go dependencies..."
go mod tidy

# Copy .env.example to .env if .env does not exist
if [ ! -f .env ]; then
    echo_green "Creating .env file from .env.example..."
    cp .env.example .env
else
    echo_green ".env file already exists."
fi

# Copy .env.example to .env.test if .env.test does not exist
if [ ! -f .env.test ]; then
    echo_green "Creating .env.test file from .env.example..."
    cp .env.example .env.test
else
    echo_green ".env.test file already exists."
fi

echo_green "Setup complete. Please review the .env file and adjust settings as necessary."
