package converter

import (
	"flag"
	"fmt"

	"os"
)

type Date struct {
	Year  int
	Month int
	Date  int
}

const (
	NEPALI_DATE = iota
	ENGLISH_DATE
)

var converterFlagSet = flag.NewFlagSet("conv", flag.ExitOnError)
var bsFlag = converterFlagSet.String("bs", "", "nepdate conv -bs <ad_date>")
var adFlag = converterFlagSet.String("ad", "", "nepdate conv -ad <bs_date>")

const ERROR_PARSING_FLAG = "Converter: Error while parsing flags %s"

func Converter() error {
	if err := converterFlagSet.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, ERROR_PARSING_FLAG, err.Error())
		os.Exit(1)
	}
	if (*adFlag == "" && *bsFlag != "") || (*adFlag != "" && *bsFlag == "") {
		if *bsFlag != "" {
			err := Convert(*bsFlag, NEPALI_DATE)
			if err != nil {
				return fmt.Errorf("Converter: %v", err)
			}
		} else {
			err := Convert(*adFlag, ENGLISH_DATE)
			if err != nil {
				return fmt.Errorf("Converter: %v", err)
			}
		}
	} else {
		return fmt.Errorf(ERROR_PARSING_FLAG, "")
	}
	return nil
}
