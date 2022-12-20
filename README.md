# dn - Daily notes command line tool

This tool is a Go implementation of [tomlockwood's dn tool](https://github.com/tomlockwood/dn).

`dn` is a simple command line tool to keep notes.

# Usage

`dn` writes a bullet-pointed string to a file with today's date in YYYY-MM-DD format in the `~/dn/` folder.

`dn -o` does the same, but the first argument is the filename. This can be used for future notes i.e. `dn -o 2030-10-01 "I died"`.

`dn -t` displays today's notes.

`dn -v` displays all files, or when an argument like `2022-12` is passed, `~/dn/2022-12*`.

`dn -s` case-insensitive search for the first argument in all notes 

`dn -S` does the same, but case-sensitive

`dn -et` edit today's notes in vim.

`dn -e` edit a note in vim for a given date. i.e. `dn -e 2022-12-12`. If no date is passed i.e. `dn -e` then a file selection prompt appears in vim.

# Example

```
$ dn "Polished this project"

$ dn -v
2022-11-30
 * Started working on the port of dn
 * Successfully ported dn
2022-12-11
 * Finished other sideproject
2022-12-12
 * Polished this project

$ dn -s po
2022-11-30:1: * Started working on the port of dn
2022-11-30:2: * Successfully ported dn           
2022-12-12:1: * Polished this project   

$ dn -S po
2022-11-30:1: * Started working on the port of dn
2022-11-30:2: * Successfully ported dn 

$ dn -v 2022-11
2022-11-30
 * Started working on the port of dn
 * Successfully ported dn

$ dn -t
2022-12-12
 * Polished this project

$ dn -o 1970-01-01 "Time starts"

$ dn -v
1970-01-01
 * Time starts
2022-11-30
 * Started working on the port of dn
 * Successfully ported dn
2022-12-11
 * Finished other sideproject
2022-12-12
 * Polished this project

$ dn -v 2022
2022-11-30
 * Started working on the port of dn
 * Successfully ported dn
2022-12-11
 * Finished other sideproject
2022-12-12
 * Polished this project
```

# Setup

Run `go install github.com/GLAD-DEV/dn@latest`