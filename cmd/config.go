package cmd

import (
	"fmt"

	"github.com/GLAD-DEV/dn/config"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

func newCmdConfig() *cobra.Command {
	root := &cobra.Command{
		Use:   "config [command]",
		Short: "Interact with dn's config.",
		Args:  cobra.ExactArgs(0),
	}

	cmdView := &cobra.Command{
		Use:   "view",
		Short: "View the configuration",
		Args:  cobra.ExactArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			conf, err := config.Load()
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			config.Print(conf)
		},
	}

	cmdAdd := &cobra.Command{
		Use:   "silentAdd [true|false]",
		Short: "If set to true, no confirmation is printed when a note is added with `dn \"This is a note\"`",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			newValue, err := argToBool(args[0])
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			conf, err := config.Load()
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			conf.SilentAdd = newValue

			err = config.Write(conf)
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			fmt.Printf("Successfully updated config:\n")
			config.Print(conf)
		},
	}

	cmdRemove := &cobra.Command{
		Use:   "silentRemove [true|false]",
		Short: "If set to true, no confirmation is printed when a note is removed.",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			newValue, err := argToBool(args[0])
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			conf, err := config.Load()
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			conf.SilentRemove = newValue

			err = config.Write(conf)
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			fmt.Printf("Successfully updated config:\n")
			config.Print(conf)
		},
	}

	root.AddCommand(cmdView, cmdAdd, cmdRemove)

	return root
}

func argToBool(arg string) (bool, error) {
	switch arg {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, fmt.Errorf("Invalid boolean value. Expected 'true' or 'false', got '%s'", arg)
	}
}
