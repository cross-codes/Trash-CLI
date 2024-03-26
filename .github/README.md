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

Here, you may replace `trashput` with any name of your choice. Move the newly created
binary into `~/.local/bin` and then execute `trashput` anywhere

---

# Usecase

This app does not have a separate retrieval functionality. It is aimed at users who
primarily use a TUI with a file manager like `Thunar` or `Dolphin` installed.
The app provides a convenient way to trash files and folders, but their
restoration should be done using a file manager.

---

Project started on: 25/03/2024

(v1.0) First functional version completed on: 26/03/2024
