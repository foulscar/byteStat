package main

import (
	"fmt"
	"os"
	"sort"
)

func (fileStat *fileStat) parse() {
	data := make([]byte, fileStat.info.Size())
	_, err := fileStat.file.Read(data)
	if err != nil {
		fmt.Println(err.Error())
		printHelp()
		os.Exit(1)
	}

	for _, byte := range data {
		if _, ok := fileStat.bytes[byte]; ok {
			fileStat.bytes[byte].count += 1
			continue
		}

		fileStat.bytes[byte] = &byteStat{count: 1}
	}
	
	sortedBytes := []byte{}

	for byte := range fileStat.bytes {
		sortedBytes = append(sortedBytes, byte)
	}
	sort.Slice(sortedBytes, func(i, j int) bool {
		iCount := fileStat.bytes[sortedBytes[i]].count
		jCount := fileStat.bytes[sortedBytes[j]].count
		return iCount > jCount	
	})

	fileStat.sortedBytes = sortedBytes
}
