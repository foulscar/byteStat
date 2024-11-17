package main

import (
	"io/fs"
	"os"
)

type fileStat struct {
	file        *os.File
	info        fs.FileInfo
	bytes       map[byte]*byteStat
	sortedBytes []byte
}

type byteStat struct {
	count int
}

func getFileStat(file *os.File, info fs.FileInfo) *fileStat {
	fileStat := &fileStat{
		file:        file,
		info:        info,
		bytes:       make(map[byte]*byteStat, info.Size()),
	}

	fileStat.parse()

	return fileStat
}
