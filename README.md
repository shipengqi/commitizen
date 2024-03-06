# commitizen

Command line utility to standardize git commit messages, golang version. Forked from [commitizen-go](https://github.com/lintingzhen/commitizen-go).

The [survey](https://github.com/AlecAivazis/survey) project is no longer maintained. Therefore, this project uses [bubbletea](https://github.com/charmbracelet/bubbletea) instead.

## Getting Started

installation with source code:

```
$ make && make install
```

or

```
$ make && ./commitizen-go install
```

commit with commitizen:

```
$ git cz
```

## Usage

```
Command line utility to standardize git commit messages.

Usage:
  commitizen
  commitizen [command]

Available Commands:
  init        Initialize this tool to git-core as git-cz.
  load        Load templates.
  help        Help about any command

Flags:
  -s, --signoff   Add a Signed-off-by trailer by the committer at the end of the commit log message.
  -a, --add       Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.
  -h, --help      help for commitizen

Use "commitizen [command] --help" for more information about a command.
```

## Configure

You can set configuration file that .git-czrc at repository root or home directory. (You can also add the extension to file, like .git-czrc.yaml) The configure file that located in repository root have a priority over the one in home directory. The format is the same as the defaultConfig string in the file [commit/defaultConfig.go]().

Type item like that:

```
package render

const DefaultCommitTemplate = `---
name: default
items:
  - name: type
    desc: "Select the type of change that you're committing:"
    type: select
    options:
      - name: feat
        desc: "A new feature"
      - name: fix
        desc: "A bug fix"
      - name: docs
        desc: "Documentation only changes"
      - name: test
        desc: "Adding missing tests"
      - name: WIP
        desc: "Work in progress"
      - name: chore
        desc: "Changes to the build process or auxiliary tools\n            and libraries such as documentation generation"
      - name: style
        desc: "Changes that do not affect the meaning of the code\n            (white-space, formatting, missing semi-colons, etc)"
      - name: refactor
        desc: "A code change that neither fixes a bug nor adds a feature"
      - name: perf
        desc: "A code change that improves performance"
      - name: revert
        desc: "Revert to a commit"
  - name: scope
    desc: "Scope. Could be anything specifying place of the commit change:"
    type: input
  - name: subject
    desc: "Subject. Concise description of the changes. Imperative, lower case and no final dot:"
    type: input
    required: true
  - name: body
    desc: "Body. Motivation for the change and contrast this with previous behavior:"
    type: textarea
  - name: footer
    desc: "Footer. Information about Breaking Changes and reference issues that this commit closes:"
    type: textarea
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"`
```

Template like that:

```
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
```