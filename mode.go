package main

import (
	"fmt"
	"time"
	"unicode"
)

type note struct {
	content string
	date    string
}

type search struct {
	needle        string
	caseSensitive bool
}

// returns exit code.
func parseModes(args []string, basePath string) int {
	if len(args) == 0 {
		fmt.Print("No arguments passed\n\n")
		showHelp()

		return exitFailure
	}

	mode := args[0]
	// Left shift args
	args = args[1:]

	switch mode {
	case "a":
		fallthrough
	case "add":
		if len(args) == 0 {
			fmt.Printf("Error: Not enough arguments passed\nUsage: dn %s {CONTENT}\n", mode)

			return exitFailure
		} else if len(args) > 1 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s {CONTENT}\n", mode)

			return exitFailure
		}

		return addNote(&note{
			content: args[0],
			date:    time.Now().Format(dateFormat),
		}, basePath)

	case "s":
		fallthrough
	case "search":
		if len(args) == 0 {
			fmt.Printf("Error: Not enough arguments passed\nUsage: dn %s {SEARCH_STR}\n", mode)

			return exitFailure
		} else if len(args) > 1 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s {SEARCH_STR}\n", mode)

			return exitFailure
		}

		caseSensitive := false
		for _, c := range args[0] {
			if unicode.IsUpper(c) {
				caseSensitive = true

				break
			}
		}

		return searchNotes(&search{
			needle:        args[0],
			caseSensitive: caseSensitive,
		}, basePath)

	case "o":
		fallthrough
	case "on":
		if len(args) < 2 {
			fmt.Printf("Error: Not enough arguments passed\nUsage: dn %s {DATE} {CONTENT}\n", mode)

			return exitFailure
		} else if len(args) > 2 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s {DATE} {CONTENT}\n", mode)

			return exitFailure
		}

		date, err := parseDate(args[0])
		if err != nil {
			fmt.Printf("Error: %s\nUsage: dn %s {DATE} {CONTENT}\n", err, mode)

			return exitFailure
		}

		return addNote(&note{
			content: args[1],
			date:    date,
		}, basePath)

	case "t":
		fallthrough
	case "today":
		if len(args) > 0 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s\n", mode)

			return exitFailure
		}

		return displayNotes(time.Now().Format(dateFormat), basePath)

	case "v":
		fallthrough
	case "view":
		switch len(args) {
		case 0:
			displayNotes("", basePath)
		case 1:
			dateSlug := args[0]
			if !isValidDateSlug(dateSlug) {
				fmt.Printf("Error: Invalid date slug\nUsage: dn %s [SLUG]\n", mode)

				return exitFailure
			}

			displayNotes(dateSlug, basePath)

		default:
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s [SLUG]\n", mode)

			return exitFailure
		}

		return exitSuccess

	case "e":
		fallthrough
	case "edit":
		if len(args) > 1 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s {DATE}\n", mode)

			return exitFailure
		}

		dateSlug := ""
		if len(args) == 1 {
			var err error
			dateSlug, err = parseDate(args[0])
			if err != nil {
				fmt.Printf("Error: %s\nUsage: dn %s {DATE}\n", err, mode)

				return exitFailure
			}
		}

		return editNote(dateSlug, basePath)

	case "et":
		fallthrough
	case "editToday":
		if len(args) > 0 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s\n", mode)

			return exitFailure
		}

		// Get today's date in ISO 8601
		return editNote(time.Now().Format(dateFormat), basePath)

	case "remove":
		fallthrough
	case "r":
		if len(args) == 0 {
			fmt.Printf("Error: Not enough arguments passed\nUsage: dn %s {DATE}\n", mode)

			return exitFailure
		} else if len(args) > 1 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s {DATE}\n", mode)

			return exitFailure
		}

		t, err := time.Parse("2006-01-02", args[0])
		if err != nil {
			t, err = time.Parse("02-01-2006", args[0])
			if err != nil {
				fmt.Println("The passed date is incorrectly formatted.\nAvailable formats are YYYY-MM-DD and DD-MM-YYYY")

				return exitFailure
			}
		}

		return remove(t.Format(dateFormat), basePath)

	default:
		// Could either be a mistyped mode or a note
		// Since modes have at most two letters, anything more than three letters is most likely a note
		if len(mode) < 3 {
			// Most likely a mistyped mode
			fmt.Printf("Error: '%s' is not a valid mode\n", mode)
			fmt.Printf("Arguments with less than three letters are interpreted as modes. To add a short note, use 'dn -a \"%s\"'\n", mode)

			return exitFailure
		}

		// Most likely a note to be added
		if len(args) > 0 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn {NOTE}\n")

			return exitFailure
		}

		status := addNote(&note{
			content: mode,
			date:    time.Now().Format(dateFormat),
		}, basePath)
		if status == exitFailure {
			return exitFailure
		}

		fmt.Printf("Added '%s' to today's note\n", mode)
		fmt.Println("If this was not intended, run `dn et` to edit today's note")

		return exitSuccess
	}
}

func parseDate(dateStr string) (string, error) {
	// Verifies that the given string matches ISO 8601
	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return "", fmt.Errorf("Invalid date passed: %w\nExpect date to be in ISO 8601 format, e.g. \"2022-01-01\"", err)
	}

	return date.Format(dateFormat), nil
}

func isValidDateSlug(dateStr string) bool {
	// Check for format 2006-01-02
	_, err := time.Parse(dateFormat, dateStr)
	if err == nil {
		return true
	}

	// Check for format 2006-01
	_, err = time.Parse("2006-01", dateStr)
	if err == nil {
		return true
	}

	// Check for format 2006
	_, err = time.Parse("2006", dateStr)

	return err == nil
}
