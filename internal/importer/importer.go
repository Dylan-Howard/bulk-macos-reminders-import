package importer

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/Dylan-Howard/reminders-import/internal/parser"
)

func Import(records []parser.Record, list string) error {
	data, err := json.Marshal(records)
	if err != nil {
		return err
	}

	// Escape single quotes for AppleScript
	safeData := strings.ReplaceAll(string(data), "'", "' & quoted form of \"'\" & '")
	script := buildAppleScript(safeData, list)

	cmd := exec.Command("osascript", "-e", script)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func buildAppleScript(jsonData string, list string) string {
	return `
use framework "Foundation"
use scripting additions

set jsonString to '` + jsonData + `'
set jsonData to current application's NSData's dataWithBytes:(jsonString's UTF8String()) length:(length of jsonString)
set parsed to current application's NSJSONSerialization's JSONObjectWithData:jsonData options:0 |error|:(missing value)

tell application "Reminders"
	if not (exists list "` + list + `") then
		set myList to make new list with properties {name:"` + list + `"}
	else
		set myList to list "` + list + `"
	end if

	repeat with itemRec in parsed
		set theTitle to (itemRec's valueForKey:"title") as text
		set theNote to (itemRec's valueForKey:"note") as text
		set theDue to (itemRec's valueForKey:"due") as text

		set ISO8601Formatter to current application's NSISO8601DateFormatter's new()
		set dueDate to ISO8601Formatter's dateFromString:theDue

		if dueDate is missing value then
			set dueDate to current date
		end if

		tell myList
			make new reminder with properties {name:theTitle, body:theNote, due date:dueDate}
		end tell
	end repeat
end tell
`
}