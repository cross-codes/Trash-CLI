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

Create `functions/dir-create.go`, and write the following:

```go
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
```

Because this is a package, declaring variables with `:=` would restrict their
scope (c.f e.g error), even if they were made exportable by capitalization.

Instead, we have to state out variable types before hand if we intend to export
them. The `init()` function is always called when we import the `functions` package
