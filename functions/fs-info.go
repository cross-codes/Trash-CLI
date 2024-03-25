package functions

import (
	"os"
	"time"
)

func DoesFileExistInTrash(trash_dir string, fname string) bool {
	_, err := os.OpenFile(trash_dir+"/info/"+fname+".trashinfo", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	return err != nil
}

func WriteFileInfo(trash_dir string, fname string, abspath string) {
	file, err := os.OpenFile(trash_dir+"/info/"+fname+".trashinfo", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic("Error in writing file information")
	}
	defer file.Close()
	currentDate := time.Now().Format("2006-01-02T15:04:05")
	content := "[Trash Info]\nPath=" + abspath + "\nDeletionDate=" + currentDate
	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		panic("Error in writing file information")
	}
}
