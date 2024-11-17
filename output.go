package main

import (
	"fmt"
	"strconv"
)

func (fileStat *fileStat) print() {
	fmt.Println("Total Bytes:", strconv.Itoa(int(fileStat.info.Size())))
	fmt.Println("Total Unique Bytes:", strconv.Itoa(len(fileStat.sortedBytes)))
	fmt.Println("\nByte    │ Stats")
	fmt.Println("────────┼────────────────")

	for _, byte := range fileStat.sortedBytes {
		count := fileStat.bytes[byte].count
		percentage := (float64(count) / float64(fileStat.info.Size())) * 100.0
		
 		fmt.Println(
			fmt.Sprintf("%08b", byte) + "│",
			count,
			strconv.FormatFloat(percentage, 'f', 3, 64) + "%",
			escapeByte(byte),
		)
	}
}

func escapeByte(b byte) (string) {
	switch b {
	case '\n':
		return `\n`
	case '\t':
		return `\t`
	case '\r':
		return `\r`
	case '\\':
		return `\\`
	case '\'':
		return `\'`
	default:
		return fmt.Sprintf("%c", b)
	}
}
