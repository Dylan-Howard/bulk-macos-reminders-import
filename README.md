# Mac OS Reminders Import

Import CSV files into macOS Reminders via CLI or right-click Quick Action.

## Introduction
This script is designed to assist students in organizing their schoolwork and managing their time more effectively within macOS. It addresses the limitation of macOS, which lacks a built-in way to import reminders or create bulk reminders. With this script, students can easily convert their assignment lists, provided in a CSV format, into reminders within the macOS Reminders app. Whether you have a list of assignments, deadlines, or tasks, simply format the CSV data, and let the script handle the rest.

## Features

- Robust CSV parsing (handles quotes, commas)
- JSON pipeline for reliability
- macOS Quick Action (right-click → import)
- Zero dependencies at runtime
- Fast Go-based processing

---

## Installation

### One-line install

```bash
curl -sL https://github.com/Dylan-Howard/reminders-import/releases/latest/download/install.sh | bash
```

## Usage
```
reminders-import parse --file tasks.csv | reminders-import import --list "My List"
```

### Supported date formats
YYYY-MM-DD
MM/DD/YYYY
RFC3339

Quick Action (Finder)

After installation:

Right-click a .csv file
Click Quick Actions
Select Import CSV to Reminders
Permissions

On first run, macOS will prompt for:

Reminders access
Automation permissions

These are required.

Development
Build
make build
Run locally
./bin/reminders-import parse --file test.csv | ./bin/reminders-import import
Project Structure
cmd/           CLI entrypoint
internal/      parser + importer
workflows/     Automator Quick Action
scripts/       install script
Release Process
Tag a release:
git tag v1.0.0
git push origin v1.0.0
GitHub Actions will:
Build binaries
Publish release artifacts
Troubleshooting
Quick Action not showing
killall Finder
Permission issues
Go to System Settings → Privacy & Security
Enable Reminders access

## CSV Data Format
To use this script, your CSV data should include the following headers: "title," "due," and "note." Your CSV should also use a comma "," as the delimiter. Create your CSV file based on the above format, ensuring that each entry under these headers contains the appropriate information for your reminders.

## Example Data
Here's a sample data set ([docs/assignments.csv](https://github.com/Dylan-Howard/AutomatedMacosReminders/blob/main/docs/assignments.csv)) that demonstrates the required format for creating reminders:

| title                                     | due              | note                          |
|-------------------------------------------|------------------|--------------------------------|
| Prerequisites: Syllabus Quiz & Course Agreement | 10/16/23 12:00 PM | www.school.com/my-course     |
| Lecture 01: Course Introduction            | 10/16/23 12:00 PM | www.school.com/my-course     |
| Lecture 02: Historical Overview Part 1     | 10/16/23 12:00 PM | www.school.com/my-course     |
| Lecture 03: Historical Overview Part 2     | 10/18/23 12:00 PM | www.school.com/my-course     |
| Lecture 04: Methods                       | 10/20/23 12:00 PM | www.school.com/my-course     |
| Lecture 05: Luther Part 1                 | 10/20/23 12:00 PM | www.school.com/my-course     |
| Lecture 06: Luther Part 2                 | 10/20/23 12:00 PM | www.school.com/my-course     |
| Reading Reflection 1                      | 10/22/23 12:00 PM | www.school.com/my-course     |
| Discussion Forum: Introductions            | 10/22/23 12:00 PM | www.school.com/my-course     |

## License
This documentation is released under the [MIT License](https://opensource.org/licenses/MIT).

## Credits and Acknowledgments
I'd like to acknowledge the assistance of [GitHub user n8henrie](https://gist.github.com/n8henrie) and their [GitHub Gist](https://gist.github.com/n8henrie/c3a5bf270b8200e33591), which provided essential help during the development of this script.
