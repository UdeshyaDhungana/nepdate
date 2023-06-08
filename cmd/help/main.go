package help

import "fmt"

func PrintHelp() {
	fmt.Println("Usage: nepdate [subcommand]")
	fmt.Println("Subcommands: ")
	fmt.Println("\t\tDisplay today's Nepali date")
	fmt.Println("conv\t\tConvert dates between A.D. and B.S.")
	fmt.Println("cal\t\tDisplay Nepali calendar")
}
