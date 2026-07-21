package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/UdeshyaDhungana/nepdate/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "date",
	Short: "Print the current date",
	Run: func(cmd *cobra.Command, args []string) {
		bsDate := internal.MustConvertAdTimeToBsDate(time.Now())
		fmt.Println(bsDate.String())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
