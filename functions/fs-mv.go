package functions

import (
	"os"
	"strconv"
	"strings"
)

func MoveFileToTrash(trash_dir string, fname string, abspath string) {
	trash_path := trash_dir + "/files/" + fname
	err := os.Rename(abspath, trash_path)
	if err != nil {
		panic("Unable to move file into trash")
	}
}

func ModifyFileName(fname string, idx int) string {
	if !strings.Contains(fname, ".") {
		return fname + "." + strconv.Itoa(idx)
	} else {
		parts := strings.Split(fname, ".")
		newParts := make([]string, 0, len(parts)+1)
		newParts = append(newParts, parts[:1]...)
		newParts = append(newParts, strconv.Itoa(idx))
		newParts = append(newParts, parts[1:]...)
		return strings.Join(newParts, ".")
	}
}
