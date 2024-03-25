package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"trashput/functions"
)

func main() {
	num_args := len(os.Args[1:])
	fmt.Println("Number of files to be trashed: ", num_args)

	// Declare directory constants and create missing directories
	functions.DirStat()

	for _, fname := range os.Args[1:] {
		abspath, err := filepath.Abs(fname)
		if err != nil {
			panic("Unable to find absolute path for " + fname)
		}
		fmt.Printf("%s\n", strings.TrimSpace(abspath))
	}
}
