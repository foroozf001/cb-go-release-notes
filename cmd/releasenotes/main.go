package main

import (
	// "fmt"
	"os"
	// "path/filepath"
	// "time"

	"github.com/go-git/go-git/v5"
	// . "github.com/go-git/go-git/v5/_examples"
	// "github.com/go-git/go-git/v5/plumbing/object"

	"vwgroup/cbit/projects/cb-go-release-notes/internal/helpers"
)

func main() {
	helpers.CheckArgs("<directory>")
	directory := os.Args[1]

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	helpers.CheckIfError(err)

	branch, _ := r.Branch("main")
	println(branch.Name)
}
