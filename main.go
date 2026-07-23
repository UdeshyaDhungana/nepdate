package main

import (
	"fmt"

	"github.com/UdeshyaDhungana/nepdate/cmd"
	"github.com/UdeshyaDhungana/nepdate/internal"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v\n", r)
		}
	}()
	internal.GetConfig()
	cmd.Execute()
}
