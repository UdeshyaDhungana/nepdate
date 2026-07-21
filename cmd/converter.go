package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/UdeshyaDhungana/nepdate/internal"
	"github.com/spf13/cobra"
)

// datePattern matches YYYY-MM-DD or YYYY/MM/DD, e.g. 2082-04-01 or 2082/04/01.
var datePattern = regexp.MustCompile(`^\d{4}[-/]\d{2}[-/]\d{2}$`)

type DateString string

func (d *DateString) String() string {
	return string(*d)
}

func (d *DateString) Set(v string) error {
	if !datePattern.MatchString(v) {
		return fmt.Errorf("invalid date %q, expected format YYYY-MM-DD or YYYY/MM/DD", v)
	}
	*d = DateString(v)
	return nil
}

func (d *DateString) Type() string {
	return "date"
}

var (
	adDate DateString
	bsDate DateString
)

var convCmd = &cobra.Command{
	Use:   "conv",
	Short: "Convert between AD and BS dates",
	Run: func(cmd *cobra.Command, _ []string) {
		if cmd.Flags().Changed("ad") {
			bsDateString := bsDate.String()
			var year, month, day int
			var err error

			year, err = strconv.Atoi(bsDateString[0:4])
			if err != nil {
				fmt.Printf("Invalid year %s\n", bsDateString)
				os.Exit(1)
			}

			month, err = strconv.Atoi(bsDateString[5:7])
			if err != nil {
				fmt.Printf("Invalid month %s\n", bsDateString)
				os.Exit(1)
			}

			day, err = strconv.Atoi(bsDateString[8:])
			if err != nil {
				fmt.Printf("Invalid day %s\n", bsDateString)
				os.Exit(1)
			}

			separator := fmt.Sprintf("%c", bsDateString[4])

			adTime := internal.MustConvertBsDateToAdTime(internal.BSDate{
				Year:  year,
				Month: internal.BSMonth(month - 1),
				Day:   day,
			})

			fmt.Printf("%d%s%02d%s%02d\n", adTime.Year(), separator, adTime.Month(), separator, adTime.Day())
		}
		if cmd.Flags().Changed("bs") {
			adDateString := adDate.String()
			var year, month, day int
			var err error

			year, err = strconv.Atoi(adDateString[0:4])
			if err != nil {
				fmt.Printf("Invalid year %s\n", adDateString)
				os.Exit(1)
			}

			month, err = strconv.Atoi(adDateString[5:7])
			if err != nil {
				fmt.Printf("Invalid month %s\n", adDateString)
				os.Exit(1)
			}

			day, err = strconv.Atoi(adDateString[8:])
			if err != nil {
				fmt.Printf("Invalid day %s\n", adDateString)
				os.Exit(1)
			}

			separator := fmt.Sprintf("%c", adDateString[4])

			bsDate := internal.MustConvertAdTimeToBsDate(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
			fmt.Printf("%d%s%02d%s%02d\n", bsDate.Year, separator, bsDate.Month+1, separator, bsDate.Day)
		}
	},
}

func init() {
	convCmd.Flags().Var(&bsDate, "ad", "BS date to convert to AD (format: YYYY-MM-DD or YYYY/MM/DD)")
	convCmd.Flags().Var(&adDate, "bs", "AD date to convert to BS (format: YYYY-MM-DD or YYYY/MM/DD)")
	convCmd.MarkFlagsOneRequired("ad", "bs")
	convCmd.MarkFlagsMutuallyExclusive("ad", "bs")
	rootCmd.AddCommand(convCmd)
}
