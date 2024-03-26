package functions

import (
	"os"
)

var (
	Trash_dir string
)

// Create and declare directories for trashing
func InitialiseTrashDirectories() {
	var err error
	home_dir, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to find home directory")
	}
	Trash_dir = home_dir + "/.local/share/Trash"
	info_dir, files_dir := Trash_dir+"/info", Trash_dir+"/files"
	_ = os.MkdirAll(Trash_dir, os.ModePerm)
	_ = os.MkdirAll(info_dir, os.ModePerm)
	_ = os.MkdirAll(files_dir, os.ModePerm)
}
