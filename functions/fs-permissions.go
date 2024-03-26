package functions

import (
	"os"
)

// Check if an item is readable (only then will the file be deletable)
func ItemIsReadable(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0444 == 0444
}
