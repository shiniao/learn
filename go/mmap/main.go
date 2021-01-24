package main

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"os"
	"path/filepath"
)

func main() {
	current, _ := os.Getwd()
	f, _ := os.OpenFile(filepath.Join(current, "mmap/file"), os.O_RDWR, 0644)
	defer f.Close()

	// use mmap read file(even bigger) to memory
	mmap, _ := mmap.Map(f, mmap.RDWR, 0)
	defer mmap.Unmap()

	fmt.Println(string(mmap))

	mmap[0] = 'X'
	mmap.Flush()
}
