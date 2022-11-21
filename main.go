package main

import (
	"os"
	"sbmf/internal"
)

func main() {
	if len(os.Args) < 2 {
		panic("which sbmf yaml file?")
	}
	file := os.Args[1]
	internal.IncreaseVersion(file)
	internal.Generate(file)
}
