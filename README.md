# dn - Daily notes command line tool

This tool is a Go implementation of [tomlockwood's dn tool](https://github.com/tomlockwood/dn).

`dn` is a simple command line tool to keep notes.

# Usage

`dn` writes a bullet-pointed string to a file with today's date in YYYY-MM-DD format.
The notes are stored in the `~/dn/` folder if `$DN_HOME` is not set.
Otherwise, the notes are stored in the folder specified by `$DN_HOME`.
`$DN_HOME` must be a valid and absolute file path.

`dn o`
does the same, but the first argument is the filename.
This can be used for future notes i.e. `dn o 2030-10-01 "I died"`.

`dn t`
displays today's notes.

`dn v`
displays all files, or when an argument like `2022-12` is passed, `~/dn/2022-12*`.

`dn s`
case-insensitive search for the first argument in all notes 

`dn S`
does the same, but case-sensitive

`dn et`
edit today's notes in vim.

`dn e`
opens the note in $EDITOR for a given date, i.e. `dn e 2022-12-12`.
If no date is passed then the editor's file selection prompt is shown. 
If the environment variable $EDITOR is not defined or empty, vim is used.

# Example

```
$ dn "Polished this project"

$ dn v
2022-11-30
 * Started working on the port of dn
 * Successfully ported dn
2022-12-11
 * Finished other sideproject
2022-12-12
 * Polished this project

$ dn s po
2022-11-30:1: * Started working on the port of dn
2022-11-30:2: * Successfully ported dn           
2022-12-12:1: * Polished this project   

$ dn S po
2022-11-30:1: * Started working on the port of dn
2022-11-30:2: * Successfully ported dn 

$ dn v 2022-11
2022-11-30
 * Started working on the port of dn
 * Successfully ported dn

$ dn t
2022-12-12
 * Polished this project

$ dn o 1970-01-01 "Time starts"

$ dn v
1970-01-01
 * Time starts
2022-11-30
 * Started working on the port of dn
 * Successfully ported dn
2022-12-11
 * Finished other sideproject
2022-12-12
 * Polished this project

$ dn v 2022
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