# Trash-CLI

This is a miniature project I am working on, to emulate moving files
into a recycling bin instead of permanently deleting them.

The project is written in Go. To understand the philosophy behind how a trash
utility works, check this [link](https://specifications.freedesktop.org/trash-spec/trashspec-latest.html)

## Initialising the project

Create `.github/`, `notes/` and `main.go`. Initialize the project using:

```bash
go mod init trashput
```

## Creating the files

Create a `main.go` file in the CWD. We only require a single
`fuctions` package, so create a directory named `functions`

Our first step is to check for the presence of a trash directory.
In systems that do not have a file manager, it is possible that
there does not exist a `~/.local/share/Trash` directory.

Create `functions/dir-creation.go`, and write the following:

```go
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
```

Because this is a package, declaring variables with `:=` would restrict their
scope (c.f e.g error), even if they were made exportable by capitalization.

Instead, we have to state out variable types before hand if we intend to export them.

Only a function named `init()` would be executed when imported,
so our function needs to be manually called later on.

Now, create the file to check for readability permissions.
According to the freedesktop specs, a file/folder must be
readable for it to be trashable.

Create `functions/fs-permissions.go`, and write the following:

```go
package functions

import (
 "os"
)

// Check if a file is readable (only then will the file be deletable)
func ItemIsReadable(filename string) bool {
 info, err := os.Stat(filename)
 if err != nil {
  return false
 }
 return info.Mode().Perm()&0444 == 0444
}
```

We want to write the logic for creating a `.trashinfo` file ;
Create `functions/fs-info.go` and write the following:

```go
package functions

import (
 "os"
 "time"
)

// Check if a .trashinfo file corresponding to fname exists
// If it does not exist, said .trashinfo file is created
func ItemExistsInTrash(trash_dir string, fname string) bool {
 _, err := os.OpenFile(trash_dir+"/info/"+fname+".trashinfo", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
 return err != nil
}

// Writes trash info content for a fname with absolute path abspath into the corresponding trash_dir
func WriteTrashInfo(trash_dir string, fname string, abspath string) {
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
```

This contains code for testing if the item exists in the trash
already. In such a case, the process of trashing should NOT
be hindered (freedesktop specs), but a unique way of naming
should be followed. It also contains code for writing file
info in the standard provided.

Now, we write the logic to trash an item and also to generate
a unique name for the same

```go
package functions

import (
 "os"
 "strconv"
 "strings"
)

// Move a given file into the corresponding trash directory
func MoveItemToTrash(trash_dir string, fname string, abspath string) {
 trash_path := trash_dir + "/files/" + fname
 err := os.Rename(abspath, trash_path)
 if err != nil {
  panic("Unable to move file into trash")
 }
}

// Attempt to create a unique indexed file/directory name
func ModifyItemName(fname string, idx int) string {
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
```

The unique fname method copies the technique used by the XFCE
file manager `Thunar`. The `idx` variable can change if even
this unique already exists.

Now, we write the main script

In `main.go`, write the following:

```go
package main

import (
 "fmt"
 "os"
 "path/filepath"
 "runtime"
 "strings"
 "trashput/functions"
)

func main() {
  osName := runtime.GOOS
  if osName != "linux" {
    panic("trashput must be used with Linux only!")
  }

 num_args := len(os.Args[1:])
 fmt.Println("Number of files to be trashed: ", num_args)

 functions.InitialiseTrashDirectories()

 for _, arg := range os.Args[1:] {
  // Obtain absolute path for the file or directory
  abspath, err := filepath.Abs(arg)
  if err != nil {
   panic("Unable to find absolute path for " + arg)
  }

  fname := filepath.Base(strings.TrimSpace(abspath)) // Determine base folder/file name
  stat := functions.ItemIsReadable(abspath)          // Check if argument is a readable folder/file

  if !stat {
   fmt.Printf("Item %s is not readable, Skipping ...", arg)
   continue
  }
  trashItem(fname, abspath)
 }
}

func trashItem(fname string, abspath string) {
 // Check is the corresponding fileinfo exists, create the trashinfo file if not
 trashStat := functions.ItemExistsInTrash(functions.Trash_dir, fname)
 if trashStat {
  // Keep trying to make a unique fname for the trashing process
  idx := 2
  for trashStat {
   fname = functions.ModifyItemName(fname, idx)
   trashStat = functions.ItemExistsInTrash(functions.Trash_dir, fname)
   idx++
  }
 }
 // Write the trashinfo and move the file into trash
 functions.WriteTrashInfo(functions.Trash_dir, fname, abspath)
 functions.MoveItemToTrash(functions.Trash_dir, fname, abspath)
}
```

The app is now complete. To add a final touch, let us display
the results of the operation in a concise table.

Create `functions/table-render.go` and define the following:

```go
package functions

import (
 "os"

 "github.com/olekukonko/tablewriter"
)

var (
 results *tablewriter.Table
)

func InitialiseTable() {
 results = tablewriter.NewWriter(os.Stdout)
 results.SetHeader([]string{"Item", "Status"})
}

func AppendSlice(slice [][]string) {
 for _, v := range slice {
  results.Append(v)
 }
}

func RenderTable() {
 results.Render()
}
```

Now rewrite the main file to call the following:

```go
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
   fname = functions.ModifyItemName(fname, idx)
   trashStat = functions.ItemExistsInTrash(functions.Trash_dir, fname)
   idx++
  }
 }
 // Write the trashinfo and move the file into trash
 functions.WriteTrashInfo(functions.Trash_dir, fname, abspath)
 functions.MoveItemToTrash(functions.Trash_dir, fname, abspath)
}
```
