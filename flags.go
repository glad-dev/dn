package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type note struct {
	content string
	date    string
}

type search struct {
	needle        string
	caseSensitive bool
}

type edit struct {
	date       string
	wantToEdit bool
}

type config struct {
	displayDate string
	addNote     note
	search      search
	edit        edit
	displayAll  bool
}

func parseFlags() config {
	addDate := flag.Bool("o", false, "Add note with passed date")
	help := flag.Bool("h", false, "Show help")
	displayToday := flag.Bool("t", false, "Display today's notes")
	displayDate := flag.Bool("v", false, "Display all notes or specific notes if date slug passed")
	searchInsensitive := flag.Bool("s", false, "Case-insensitive search in all notes")
	searchSensitive := flag.Bool("S", false, "Case-sensitive search in all notes")
	editToday := flag.Bool("et", false, "Edit today's note in vim")
	editDate := flag.Bool("e", false, "Edit note with passed date in vim. If no date is passed, vim's file selection is opened")

	flag.Parse()
	args := flag.Args()

	// Check if more than one flag is set
	allFlags := []bool{*addDate, *help, *displayToday, *displayDate, *editToday, *editDate, *searchInsensitive, *searchSensitive}
	flagSet := false
	for _, item := range allFlags {
		if item {
			if flagSet {
				fmt.Printf("Mutiple flags used\n\n")
				showHelp()
				os.Exit(1)
			}

			flagSet = true
		}
	}

	if len(args) == 0 && !flagSet {
		fmt.Printf("No arguments passed\n\n")
		showHelp()
		os.Exit(1)
	}

	conf := config{
		displayAll:  false,
		displayDate: "",
		addNote: note{
			content: "",
			date:    "",
		},
		search: search{
			needle:        "",
			caseSensitive: false,
		},
		edit: edit{
			wantToEdit: false,
			date:       "",
		},
	}

	if *help {
		// Show help
		showHelp()
		os.Exit(0)
	}

	if !flagSet {
		// Add note with today's date
		if len(args) > 1 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn {CONTENT}\n")
			os.Exit(1)
		}

		conf.addNote = note{
			content: args[0],
			date:    time.Now().Format(dateFormat),
		}
	}

	if *addDate {
		// Add note with date
		if len(args) < 2 {
			fmt.Print("Error: Not enough arguments passed\nUsage: dn -o {DATE} {CONTENT}\n")
			os.Exit(1)
		} else if len(args) > 2 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn -o {DATE} {CONTENT}\n")
			os.Exit(1)
		}

		date, err := parseDate(args[0])
		if err != nil {
			fmt.Printf("Error: %s\nUsage: dn -o {DATE} {CONTENT}\n", err)
			os.Exit(1)
		}

		conf.addNote = note{
			content: args[1],
			date:    date,
		}
	}

	if *searchInsensitive {
		if len(args) < 1 {
			fmt.Print("Error: Not enough arguments passed\nUsage: dn -s {SEARCH_STR}\n")
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn -s {SEARCH_STR}\n")
			os.Exit(1)
		}

		conf.search = search{
			needle:        strings.ToLower(args[0]),
			caseSensitive: false,
		}
	}

	if *searchSensitive {
		if len(args) < 1 {
			fmt.Print("Error: Not enough arguments passed\nUsage: dn -S {SEARCH_STR}\n")
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn -S {SEARCH_STR}\n")
			os.Exit(1)
		}

		conf.search = search{
			needle:        args[0],
			caseSensitive: true,
		}
	}

	if *displayToday {
		// Display today's notes
		if len(args) > 0 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn -t\n")
			os.Exit(1)
		}

		conf.displayDate = time.Now().Format(dateFormat)
	}

	if *displayDate {
		// Display all notes if no additional parameter is given.
		// Otherwise, display all notes that match the RegEx
		switch len(args) {
		case 0:
			conf.displayAll = true
		case 1:
			date := args[0]
			if !isValidDateSlug(date) {
				fmt.Print("Error: Invalid date slug\nUsage: dn -t [SLUG]\n")
				os.Exit(1)
			}

			conf.displayDate = date
		default:
			fmt.Print("Error: Too many arguments passed\nUsage: dn -t [SLUG]\n")
			os.Exit(1)
		}
	}

	if *editToday {
		if len(args) > 0 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn -et\n")
			os.Exit(1)
		}

		// Get today's date in ISO 8601
		conf.edit = edit{
			wantToEdit: true,
			date:       time.Now().Format(dateFormat),
		}
	}

	if *editDate {
		date := ""
		if len(args) > 1 {
			fmt.Print("Error: Too many arguments passed\nUsage: dn -et {DATE}\n")
			os.Exit(1)
		} else if len(args) == 1 {
			var err error
			date, err = parseDate(args[0])
			if err != nil {
				fmt.Printf("Error: %s\nUsage: dn -eo {DATE}\n", err)
				os.Exit(1)
			}
		}

		conf.edit = edit{
			wantToEdit: true,
			date:       date,
		}
	}

	return conf
}

func showHelp() {
	fmt.Printf("usage: dn [-h] [-s] [-o] [-t] [-v] [-e] [-et] {content}\n\n")

	fmt.Println("optional arguments:")
	fmt.Println("  -h\t\tShow this help message and quit")
	fmt.Println("  -s\t\tSearch for a string in all notes")
	fmt.Println("  -o [DATE]\tAdd note dated at [DATE]")
	fmt.Println("  -t\t\tDisplay today's notes")
	fmt.Println("  -v [DATE]\tDisplays all notes if no date is passed, otherwise display the notes that match the passed date slug")
	fmt.Println("  -e [DATE]\tIf no date is passed, vim's file selection is opened. Otherwise the corresponding note is opened in vim")
	fmt.Println("  -et\t\tEdit today's note in vim")
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
