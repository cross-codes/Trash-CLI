<div align="center">
<h1> 🗑️Trash-CLI </h1>

A command-line utility to trash files safely on any UNIX-like operating system,
compliant with the FreeDesktop.org Trash specification, written in Go

Current version : 1.0

For a brief overview on the process of creating the app, check out
`📁notes/process.md`

</div>

---

# Installation

To install the app, first ensure you have `go` installed and on your path.
Then, clone the repo and build the binary:

```zsh
git clone https://github.com/cross-codes/Trash-CLI.git
go build -ldflags="-s -w" -o trashput
```

Here, you may replace `trashput` with any name of your choice
which will now be the command used to execute the application). Move the newly created
binary into a location included in PATH such as `~/.local/bin` and then
execute `trashput` anywhere:

```bash
trashput [FILE/DIR]...
```

The arguments can contain globbing patterns

---

# Usecase

This app is designed for users who primarily use a TUI with a file manager like
`Thunar` or `Dolphin` installed. It provides a convenient way to trash files and
folders. However, it does not include a separate restoration functionality.
The decision was made to keep the scope focused on trashing files efficiently,
while leaving the restoration process to be handled through the user's file manager.

This specialization allows for a lightweight and streamlined user experience,
especially for those accustomed to managing files primarily through graphical interfaces

---

Project started on: 25/03/2024

(v1.0) First functional version completed on: 26/03/2024
