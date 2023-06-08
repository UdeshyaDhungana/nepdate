package main

import (
	"fmt"
	"nepdate/cmd/calendar"
	"nepdate/cmd/converter"
	"nepdate/cmd/help"
	"nepdate/cmd/nepdate"
	"os"
)

// subcommands and their flags
const (
	SUBCOMMAND_HELP = "help"
	SUBCOMMAND_CONV = "conv"
	SUBCOMMAND_CAL  = "cal"
)

const WRONG_SUBCOMMAND_ERROR_MSG_FORMAT = "nepdate: invalid subcommand: %s\n"

func main() {
	if len(os.Args) < 2 {
		err := nepdate.NepDate()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err.Error())
			os.Exit(1)
		}
	} else {
		subCommand := os.Args[1]
		switch subCommand {
		case SUBCOMMAND_CONV:
			converter.Converter()
		case SUBCOMMAND_CAL:
			calendar.Calendar()
		case SUBCOMMAND_HELP:
			help.PrintHelp()
		default:
			fmt.Fprintf(os.Stderr, WRONG_SUBCOMMAND_ERROR_MSG_FORMAT, subCommand)
			help.PrintHelp()
			os.Exit(1)
		}
	}
}
