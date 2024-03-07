# commitizen

Command line utility to standardize git commit messages, golang version. Forked from [commitizen-go](https://github.com/lintingzhen/commitizen-go).

The [survey](https://github.com/AlecAivazis/survey) project is no longer maintained. Therefore, this project uses [bubbletea](https://github.com/charmbracelet/bubbletea) instead.

## Getting Started

```
Command line utility to standardize git commit messages.

Usage:
  commitizen
  commitizen [command]

Available Commands:
  init        Initialize this tool to git-core as git-cz.
  help        Help about any command

Flags:
  -s, --signoff   add a Signed-off-by trailer by the committer at the end of the commit log message.
  -a, --add       tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.
      --dry-run   you can use the --dry-run flag to preview the message that would be committed, without really submitting it.
  -h, --help      help for commitizen

Use "commitizen [command] --help" for more information about a command.
```

commit with commitizen:

```
$ git cz
```

## Installing commitizen 

### From the Binary Releases

### From Source

You must have a working Go environment:

```
$ git clone https://github.com/shipengqi/commitizen.git
$ cd commitizen
$ make && make install
```

Or:

```
$ make && ./commitizen-go install
```

## Configuration

You can set configuration file that `.git-czrc` at repository root or home directory. The configuration file that located in repository root have a priority over the one in home directory. The format is the same as the following:

```yaml
name: my-default
default: true  # (optional) If true, this template will be used as the default template, note that there can only be one default template       
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

Commit message `format`:

```
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
```