package parser

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"
	"time"
)

type Record struct {
	Title string `json:"title"`
	Due   string `json:"due"`
	Note  string `json:"note"`
}

var expectedHeaders = []string{"title", "due", "note"}

func ParseCSV(r io.Reader, strict bool) ([]Record, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, errors.New("no data rows")
	}

	if !equal(rows[0], expectedHeaders) {
		return nil, errors.New("invalid headers: expected title,due,note")
	}

	var out []Record

	for i, row := range rows[1:] {
		if len(row) != 3 {
			if strict {
				return nil, errors.New("invalid column count on row")
			}
			continue
		}

		title := strings.TrimSpace(row[0])
		dueRaw := strings.TrimSpace(row[1])
		note := strings.TrimSpace(row[2])

		dueISO, err := normalizeDate(dueRaw)
		if err != nil {
			if strict {
				return nil, err
			}
			continue
		}

		out = append(out, Record{
			Title: title,
			Due:   dueISO,
			Note:  note,
		})

		_ = i
	}

	return out, nil
}

func normalizeDate(input string) (string, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"1/2/2006",
		"1/2/06",
	}

	for _, l := range layouts {
		if t, err := time.Parse(l, input); err == nil {
			return t.Format(time.RFC3339), nil
		}
	}

	return "", errors.New("invalid date: " + input)
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if strings.ToLower(strings.TrimSpace(a[i])) != b[i] {
			return false
		}
	}
	return true
}