// Package cmd defines dn's command-line interface.
package cmd

import (
	"strings"

	"github.com/GLAD-DEV/dn/constants"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "dn",
		Version: constants.Version,
	}

	cmdAdd := newCmdAdd()
	cmdEdit := newCmdEdit()
	cmdEditToday := newCmdEditToday()
	cmdOn := newCmdOn()
	cmdRemove := newCmdRemove()
	cmdSearch := newCmdSearch()
	cmdToday := newCmdToday()
	cmdView := newCmdView()
	cmdConfig := newCmdConfig()

	rootCmd.AddCommand(cmdAdd, cmdEdit, cmdEditToday, cmdOn, cmdRemove, cmdSearch, cmdToday, cmdView, cmdConfig)

	showAliasesInUsage(rootCmd)

	return rootCmd
}

// showAliasesInUsage lists subcommands as "add (a)" instead of "add" in the help output.
// Subcommands inherit the usage template, so this also applies to e.g. `dn config --help`.
func showAliasesInUsage(rootCmd *cobra.Command) {
	cobra.AddTemplateFunc("nameWithAliases", nameWithAliases)
	cobra.AddTemplateFunc("aliasPadding", func(c *cobra.Command) int {
		padding := 0
		for _, sibling := range c.Parent().Commands() {
			padding = max(padding, len(nameWithAliases(sibling)))
		}

		return padding
	})

	rootCmd.SetUsageTemplate(strings.ReplaceAll(
		rootCmd.UsageTemplate(),
		"{{rpad .Name .NamePadding }}",
		"{{rpad (nameWithAliases .) (aliasPadding .)}}",
	))
}

func nameWithAliases(c *cobra.Command) string {
	if len(c.Aliases) == 0 {
		return c.Name()
	}

	return c.Name() + " (" + strings.Join(c.Aliases, ", ") + ")"
}
