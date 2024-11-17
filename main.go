package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		printHelp()
		return
	}

	fileNameAbs, _ := filepath.Abs(os.Args[1])

	info, err := os.Stat(fileNameAbs)
	if err != nil {
		fmt.Println("File does not exist or you do not have the permissions")
		printHelp()
		os.Exit(1)
	}
	if info.IsDir() {
		fmt.Println("'", os.Args[1], "' is a directory")
		printHelp()
		os.Exit(1)
	}

	file, err := os.Open(fileNameAbs)
	if err != nil {
		fmt.Println(err.Error())
		printHelp()
		os.Exit(1)
	}

	fileStat := getFileStat(file, info)
	fileStat.print()
}
