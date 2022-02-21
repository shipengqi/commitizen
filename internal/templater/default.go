package templater

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
      - name: style
        desc: "Changes that do not affect the meaning of the code\n  (white-space, formatting, missing semi-colons, etc)"
      - name: refactor
        desc: 'A code change that neither fixes a bug nor adds a feature'
      - name: perf
        desc: 'A code change that improves performance'
      - name: test
        desc: 'Adding missing tests'
      - name: ci
        desc: "A code change for CI/CD"
      - name: revert
        desc: 'Revert to a commit'
      - name: WIP
        desc: 'Work in progress'
    required: true
  - name: scope
    desc: 'Scope. Could be anything specifying place of the commit change:'
    type: input
  - name: subject
    desc: 'Subject. Concise description of the changes. Imperative, lower case and no final dot:'
    type: input
    required: true
  - name: body
    desc: 'Body. Motivation for the change and contrast this with previous behavior:'
    type: multiline
  - name: footer
    desc: 'Footer. Information about Breaking Changes and reference issues that this commit closes:'
    type: multiline
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"`

