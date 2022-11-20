package main

import (
	"os"
	"sbmf/internal"
)

func main() {
	if len(os.Args) < 2 {
		panic("which sbmf yaml file?")
	}
	internal.Generate(os.Args[1])
}
