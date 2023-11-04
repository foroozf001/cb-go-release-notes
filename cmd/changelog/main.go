package main

import (
	utils "github.com/foroozf001/cb-go-release-notes/internal/utils"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	// "github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	utils.CheckArgs("<new version>", "<previous version>", "<directory>")
	newVersion := os.Args[1]
	oldVersion := os.Args[2]
	directory := os.Args[3]

	// Opens an already existing repository
	r, err := git.PlainOpen(directory)
	utils.CheckIfError(err)

	// ...retrieves all tag references and determines commit window
	var since, until time.Time
	tags, err := r.TagObjects()
	utils.CheckIfError(err)
	err = tags.ForEach(func(t *object.Tag) error {
		obj, err := r.CommitObject(t.Target)
		utils.CheckIfError(err)

		if t.Name == newVersion {
			until = obj.Committer.When
		} else if t.Name == oldVersion {
			since = obj.Committer.When
		}
		return nil
	})
	utils.CheckIfError(err)

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	utils.CheckIfError(err)

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	utils.CheckIfError(err)

	// ... just iterates over the commits
	commits := make(map[string]string)
	err = cIter.ForEach(func(c *object.Commit) error {
		commits[c.Hash.String()] = strings.Split(c.Message, "\n")[0]
		return nil
	})
	utils.CheckIfError(err)

	release := utils.Changelog{
		Version: newVersion,
		Commits: commits,
		Date:    time.Now().Format("January 2, 2006"),
	}
	release.Template()
}
