#!/usr/bin/env bash
set -e

REPO="Dylan-Howard/bulk-macos-reminders-import"
BIN_NAME="reminders-import"

# Determine architecture
ARCH=$(uname -m)
if [[ "$ARCH" == "arm64" ]]; then
  URL="https://github.com/$REPO/releases/latest/download/$BIN_NAME-darwin-arm64"
else
  URL="https://github.com/$REPO/releases/latest/download/$BIN_NAME-darwin-amd64"
fi

echo "Downloading binary from $URL..."
if ! curl -L "$URL" -o "$BIN_NAME"; then
    echo "Error: Failed to download binary. Please check if the release exists."
    exit 1
fi
chmod +x "$BIN_NAME"

echo "Installing binary to /usr/local/bin/..."
if [[ -w "/usr/local/bin" ]]; then
    mv "$BIN_NAME" "/usr/local/bin/"
else
    echo "Requires sudo to install to /usr/local/bin"
    sudo mv "$BIN_NAME" "/usr/local/bin/"
fi

echo "Installing Quick Action..."
SERVICES_DIR="$HOME/Library/Services"
mkdir -p "$SERVICES_DIR"

WORKFLOW_NAME="Import CSV to Reminders.workflow"

# If we are in the repo, copy the full workflow bundle
if [[ -d "workflows/$WORKFLOW_NAME" ]]; then
    echo "Copying Quick Action from workflows directory..."
    cp -R "workflows/$WORKFLOW_NAME" "$SERVICES_DIR/"
else
    echo "Creating Quick Action from template..."
    # Fallback to creating a basic version if the bundle isn't available locally
    mkdir -p "$SERVICES_DIR/$WORKFLOW_NAME/Contents"
    cat > "$SERVICES_DIR/$WORKFLOW_NAME/Contents/document.wflow" << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>actions</key>
	<array>
		<dict>
			<key>action</key>
			<dict>
				<key>ActionParameters</key>
				<dict>
					<key>shell</key>
					<string>/bin/bash</string>
					<key>script</key>
					<string>
for f in "$@"
do
  /usr/local/bin/reminders-import parse --file "$f" | /usr/local/bin/reminders-import import
done
					</string>
				</dict>
				<key>BundleIdentifier</key>
				<string>com.apple.RunShellScript</string>
			</dict>
		</dict>
	</array>
	<key>workflowMetaData</key>
	<dict>
		<key>workflowTypeIdentifier</key>
		<string>com.apple.Automator.servicesMenu</string>
	</dict>
</dict>
</plist>
EOF
fi

echo "Installed successfully."
echo "Note: You may need to enable the Quick Action in System Settings > Keyboard > Keyboard Shortcuts > Services."
