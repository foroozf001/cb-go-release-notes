package utils

import (
	"fmt"
	"os"
	"text/template"
)

type Changelog struct {
	Version string
	Commits map[string]string
	Date    string
}

func (c Changelog) Template() {
	// Files are provided as a slice of strings.
	changelogTemplate, err := os.ReadFile("Changelog.md.gotmpl")
	if err != nil {
		_ = fmt.Errorf("could not read changelog template file %s", err)
	}
	t := template.Must(template.New("changelog").Parse(string(changelogTemplate)))
	// create a new file
	file, err := os.Create(fmt.Sprintf("changelog/Changelog-%s.md", c.Version))
	if err != nil {
		_ = fmt.Errorf("could not create changelog file %s", err)
	}
	defer file.Close()

	err = t.Execute(file, c)

	if err != nil {
		_ = fmt.Errorf("executing template: %s", err)
	}
}
