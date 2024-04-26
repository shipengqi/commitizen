package config

const DefaultCommitTemplateV1 = `---
name: default
default: true
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
        desc: "Adding missing or correcting existing tests"
      - name: chore
        desc: "Changes to the build process or auxiliary tools and\n               libraries such as documentation generation"
      - name: style
        desc: "Changes that do not affect the meaning of the code\n               (white-space, formatting, missing semi-colons, etc)"
      - name: refactor
        desc: "A code change that neither fixes a bug nor adds a feature"
      - name: perf
        desc: "A code change that improves performance"
      - name: revert
        desc: "Reverts a previous commit"
  - name: scope
    desc: "Scope. What is the scope of this change? (class or file name):"
    type: input
  - name: subject
    desc: "Subject. Write a short and imperative summary of the code change (lower case and no period):"
    type: input
    required: true
  - name: body
    desc: "Body. Provide additional contextual information about the code changes:"
    type: textarea
  - name: footer
    desc: "Footer. Information about Breaking Changes and reference issues that this commit closes:"
    type: textarea
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"`

const DefaultCommitTemplate = `---
version: v2
name: default
default: true
items:
  - name: type
    label: "Select the type of change that you're committing:"
    type: list
    options:
      - value: feat
        key: "feat:      A new feature"
      - value: fix
        key: "fix:       A bug fix"
      - value: docs
        key: "docs:      Documentation only changes"
      - value: test
        key: "test:      Adding missing or correcting existing tests"
      - value: chore
        key: "chore:     Changes to the build process or auxiliary tools and\n             libraries such as documentation generation"
      - value: style
        key: "style:     Changes that do not affect the meaning of the code\n             (white-space, formatting, missing semi-colons, etc)"
      - value: refactor
        key: "refactor:  A code change that neither fixes a bug nor adds a feature"
      - value: perf
        key: "perf:      A code change that improves performance"
      - value: revert
        key: "revert:    Reverts a previous commit"
  - name: scope
    label: "Scope. What is the scope of this change? (class or file name):"
    type: string
  - name: subject
    label: "Subject. Write a short and imperative summary of the code change (lower case and no period):"
    type: string
    required: true
  - name: body
    label: "Body. Provide additional contextual information about the code changes:"
    type: text
  - name: footer
    label: "Footer. Information about Breaking Changes and reference issues that this commit closes:"
    type: text
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"`
