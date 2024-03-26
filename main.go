package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"trashput/functions"
)

var status [][]string

func main() {
	osName := runtime.GOOS
	if osName != "linux" {
		panic("trashput must be used with Linux only!")
	}

	num_args := len(os.Args[1:])
	if num_args == 0 {
		panic("trashput: missing operand")
	}

	functions.InitialiseTrashDirectories()

	for _, arg := range os.Args[1:] {
		// Obtain absolute path for the file or directory
		abspath, err := filepath.Abs(arg)
		if err != nil {
			panic("Unable to construct absolute path for " + arg)
		}

		fname := filepath.Base(strings.TrimSpace(abspath)) // Determine base folder/file name
		stat := functions.ItemIsReadable(abspath)          // Check if argument is a readable folder/file

		if !stat {
			fmt.Printf("[WARN] Item %s is not readable. Skipping ...\n", arg)
			status = append(status, []string{fname, "Skipped"})
			continue
		}
		trashItem(fname, abspath)
		status = append(status, []string{fname, "Removed"})
	}
	if num_args > 0 {
		functions.InitialiseTable()
		functions.AppendSlice(status)
		functions.RenderTable()
	}
}

func trashItem(fname string, abspath string) {
	// Check is the corresponding fileinfo exists, create the trashinfo file if not
	trashStat := functions.ItemExistsInTrash(functions.Trash_dir, fname)
	if trashStat {
		// Keep trying to make a unique fname for the trashing process
		idx := 2
		for trashStat {
			fname = functions.ModifyItemBaseName(fname, idx)
			trashStat = functions.ItemExistsInTrash(functions.Trash_dir, fname)
			idx++
		}
	}
	// Write the trashinfo and move the item into trash
	functions.WriteTrashInfo(functions.Trash_dir, fname, abspath)
	functions.MoveItemToTrash(functions.Trash_dir, fname, abspath)
}
