package main

import (
	"github.com/UdeshyaDhungana/nepdate/cmd"
	"github.com/UdeshyaDhungana/nepdate/internal"
)

func main() {
	internal.GetConfig()
	cmd.Execute()
}
