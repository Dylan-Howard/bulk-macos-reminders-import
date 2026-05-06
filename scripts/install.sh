#!/usr/bin/env bash
set -e

REPO="yourname/reminders-import"
BIN_NAME="reminders-import"

ARCH=$(uname -m)
if [[ "$ARCH" == "arm64" ]]; then
  URL="https://github.com/$REPO/releases/latest/download/$BIN_NAME-darwin-arm64"
else
  URL="https://github.com/$REPO/releases/latest/download/$BIN_NAME-darwin-amd64"
fi

echo "Downloading binary..."
curl -L $URL -o $BIN_NAME
chmod +x $BIN_NAME

echo "Installing binary..."
mv $BIN_NAME /usr/local/bin/

echo "Installing Quick Action..."
mkdir -p "$HOME/Library/Services"

cat > "$HOME/Library/Services/Import CSV to Reminders.workflow/Contents/document.wflow" << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<workflow>
  <action>
    <shellScript>
for f in "$@"
do
  /usr/local/bin/reminders-import parse --file "$f" | /usr/local/bin/reminders-import import
done
    </shellScript>
  </action>
</workflow>
EOF

echo "Installed successfully."