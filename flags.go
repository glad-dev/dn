package main

import (
	"flag"
	"fmt"
)

const (
	dateFormat  = "2006-01-02"
	version     = "1.5.0"
	exitSuccess = 0
	exitFailure = 1
)

func parseArguments() int {
	help := flag.Bool("h", false, "Show help")
	helpLong := flag.Bool("help", false, "Show help")
	versionFlag := flag.Bool("V", false, "Displays the tool's version")
	versionFlagLong := flag.Bool("version", false, "Displays the tool's version")

	flag.Parse()

	if *versionFlag || *versionFlagLong {
		fmt.Printf("dn %s\n", version)

		return exitSuccess
	} else if *help || *helpLong {
		showHelp()

		return exitSuccess
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Print("No arguments passed\n\n")
		showHelp()

		return exitFailure
	}

	return parseModes(args)
}

func showHelp() {
	fmt.Print("usage: dn [-h] [-V] {add, search, sensitiveSearch, ...}\n\n")

	fmt.Println("optional arguments:")
	fmt.Println("  -h\t\tShow this help message and quit")
	fmt.Print("  -V\t\tDisplays the tool's version\n\n")

	fmt.Println("modes:")
	fmt.Println(
		"  add (a)" +
			"\t\tAdds note with today's date. Can be omitted if note is longer than three chars.",
	)
	fmt.Println(
		"  search (s)" +
			"\t\tIf the argument contains at least one capital letter, the search is case-sensitive." +
			"\n\t\t\tOtherwise, it is case-insensitive.",
	)
	fmt.Println(
		"  on (o)" +
			"\t\tAdd note dated at passed date.",
	)
	fmt.Println(
		"  today (t)" +
			"\t\tDisplay today's notes.",
	)
	fmt.Println(
		"  view (v)" +
			"\t\tIf no date is passed, displays all notes. Otherwise displays notes matching passed date slug.",
	)
	fmt.Println(
		"  edit (e)" +
			"\t\tIf no date is passed, $EDITOR's file selection is opened. Otherwise the corresponding note" +
			"\n\t\t\tis opened in $EDITOR.",
	)
	fmt.Println(
		"  editToday (et)" +
			"\tEdit today's note in $EDITOR.",
	)
	fmt.Println(
		"  remove (r)" +
			"\t\tRemoves the notes from the passed date.",
	)
}
