package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Dylan-Howard/reminders-import/internal/importer"
	"github.com/Dylan-Howard/reminders-import/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: reminders-import [parse|import]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "parse":
		parseCmd(os.Args[2:])
	case "import":
		importCmd(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func parseCmd(args []string) {
	fs := flag.NewFlagSet("parse", flag.ExitOnError)
	input := fs.String("file", "", "CSV file path")
	strict := fs.Bool("strict", false, "Fail on invalid rows")
	fs.Parse(args)

	if *input == "" {
		fmt.Println("Missing --file")
		os.Exit(1)
	}

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	records, err := parser.ParseCSV(f, *strict)
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(records)
}

func importCmd(args []string) {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	list := fs.String("list", "Scripted Reminders", "Reminders list name")
	fs.Parse(args)

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var records []parser.Record
	if err := json.Unmarshal(input, &records); err != nil {
		panic(err)
	}

	if err := importer.Import(records, *list); err != nil {
		panic(err)
	}
}