package functions

import (
	"os"
)

var (
  Home_dir  string
	Trash_dir string
	Info_dir  string
	Files_dir string
)

func DirStat() {
  var err error
	Home_dir, err = os.UserHomeDir()
	if err != nil {
		panic("Unable to find home directory")
	}
	Trash_dir = Home_dir + "/.local/share/Trash"
	Info_dir, Files_dir = Trash_dir+"/info", Trash_dir+"/files"
	_ = os.MkdirAll(Trash_dir, os.ModePerm)
	_ = os.MkdirAll(Info_dir, os.ModePerm)
	_ = os.MkdirAll(Files_dir, os.ModePerm)
}
