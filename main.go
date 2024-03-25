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

	functions.DirStat()

	for _, arg := range os.Args[1:] {
		abspath, err := filepath.Abs(arg)
		if err != nil {
			panic("Unable to find absolute path for " + arg)
		}

		fname := filepath.Base(strings.TrimSpace(abspath))
		stat := functions.FileIsReadable(strings.TrimSpace(abspath))

		if !stat {
			fmt.Printf("Filename %s is not readable, Skipping ...", arg)
			continue
		}

		trashStat := functions.DoesFileExistInTrash(functions.Trash_dir, fname)
		if !trashStat {
			functions.WriteFileInfo(functions.Trash_dir, fname, abspath)
		} else {
			fmt.Println("Hello")
		}
	}
}
