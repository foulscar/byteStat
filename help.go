package main

import (
	"fmt"
)

const helpMessage = `
Usage: byteStat [OPTIONS] <INPUTFILE>

Options:
  help        Show this help message and exit

Examples:
  byteStat help
  byteStat file.txt
  byteStat file.txt > file.txt.stat`

func printHelp() {
	fmt.Println(helpMessage)
}
