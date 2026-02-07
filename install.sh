# Builds the project and sets up executable.
set -e

echo "Building..."
go build -o terminal-notes

echo "Installing Executable..."
mv terminal-notes /usr/local/bin

echo "Done."