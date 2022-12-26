package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type note struct {
	content string
	date    string
}

type search struct {
	needle        string
	caseSensitive bool
}

// returns exit code
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

		addNote(&note{
			content: args[0],
			date:    time.Now().Format(dateFormat),
		}, basePath)

		return exitSuccess
	case "s":
		fallthrough
	case "search":
		if len(args) == 0 {
			fmt.Printf("Error: Not enough arguments passed\nUsage: dn %s {SEARCH_STR}\n", mode)
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s {SEARCH_STR}\n", mode)
			os.Exit(1)
		}

		searchNotes(&search{
			needle:        strings.ToLower(args[0]),
			caseSensitive: false,
		}, basePath)

		return exitSuccess
	case "S":
		fallthrough
	case "sensitiveSearch":
		if len(args) == 0 {
			fmt.Printf("Error: Not enough arguments passed\nUsage: dn %s {SEARCH_STR}\n", mode)
			return exitFailure
		} else if len(args) > 1 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s {SEARCH_STR}\n", mode)
			return exitFailure
		}

		searchNotes(&search{
			needle:        args[0],
			caseSensitive: true,
		}, basePath)

		return exitSuccess
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

		addNote(&note{
			content: args[1],
			date:    date,
		}, basePath)

		return exitSuccess
	case "t":
		fallthrough
	case "today":
		if len(args) > 0 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s\n", mode)
			return exitFailure
		}

		displayNotes(time.Now().Format(dateFormat), basePath)

		return exitSuccess
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

		editNote(dateSlug, basePath)
		return exitSuccess
	case "et":
		fallthrough
	case "editToday":
		if len(args) > 0 {
			fmt.Printf("Error: Too many arguments passed\nUsage: dn %s\n", mode)
			return exitFailure
		}

		// Get today's date in ISO 8601
		editNote(time.Now().Format(dateFormat), basePath)
		return exitSuccess
	default:
		// Could either be a mistyped mode or a note
		// Since modes have at most two letters, anything more than three letters is most likely a note
		if len(mode) < 3 {
			// Most likely a mistyped mode
			fmt.Printf("Unknown mode passed %s\n\n", mode)
			showHelp()
			return exitFailure
		}

		// Most likely a note to be added
		if len(args) > 0 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn {NOTE}\n")
			return exitFailure
		}

		addNote(&note{
			content: mode,
			date:    time.Now().Format(dateFormat),
		}, basePath)

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
	if err == nil {
		return true
	}

	return false
}
