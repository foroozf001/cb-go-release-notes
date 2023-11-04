# Release notes

Every new release to codebeamer or extensions introduces changes. It's a common practice to document the changes in documents called release notes or changelogs.

Often the words release notes and changelogs are used interchangeably because there's only a slight difference between them. Where release notes are non-technical and target end-users, changelogs are technical and target developers.

This project demonstrates how to automatically generate release documents for any existing Git repository. These documents always contain the artifacts and commit history.

## Getting started

To start using the `changelog` binary, the target Git repository must have release commits marked with `git tags`. For example:

```bash
$ git log --oneline
e1d3042 (HEAD -> main, origin/main) Added date to template
b45f6e7 Added 0.1.4 changelog
89197bf (tag: 0.1.4) Create version 0.1.4
26929f8 Refactor template
9c5efeb Refactor name to changelog
7291909 Added draft release notes
743be2a Added tags hash map
84b520a (tag: 0.1.3) Added stub for tags
```

### Build for Linux

```bash
make build GOOS=linux GOARCH=amd64
```

### Build for OSX

```bash
make build GOOS=darwin GOARCH=arm64
```

### Generate changelog

```bash
make changelog NEW_VERSION=0.1.5 OLD_VERSION=0.1.4
```

The new changelog file is stored under `changelog/`.

## Help

```bash
$ make help

Usage:
  make <target>
  help             Displays help

Development
  fmt              Formats the source code
  vet              Validates the source code
  run              Runs the source code
  build            Builds the binary
  changelog        Generates changelogs
```
