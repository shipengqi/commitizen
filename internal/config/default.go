package config

// DefaultCommitTemplate is the default commit template.
const DefaultCommitTemplate = `---
name: default
default: true
groups:
  - name: hasbreaking
    depends_on:
      and_conditions:
      - parameter_name: page2.isbreaking
        value_equals: true
  - name: nobreaking
    depends_on:
      and_conditions:
      - parameter_name: page2.isbreaking
        value_equals: false
items:
  - name: type
    group: page1
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
        key: "chore:     Changes to the build process or auxiliary tools and libraries such as documentation generation"
      - value: style
        key: "style:     Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"
      - value: refactor
        key: "refactor:  A code change that neither fixes a bug nor adds a feature"
      - value: perf
        key: "perf:      A code change that improves performance"
      - value: revert
        key: "revert:    Reverts a previous commit"
  - name: scope
    group: page2
    label: "Scope. What is the scope of this change? (class or file name):"
    type: string
    trim: true
  - name: subject
    group: page2
    label: "Subject. Write a short and imperative summary of the code change (lower case and no period):"
    type: string
    required: true
    trim: true
  - name: isbreaking
    group: page2
    label: "Are there any breaking changes?"
    type: boolean
  - name: hasbreakingbody
    group: hasbreaking
    label: "A BREAKING CHANGE commit requires a body. Provide additional contextual information about the code changes:"
    type: text
    required: true
  - name: nobreakingbody
    group: nobreaking
    label: "Body. Provide additional contextual information about the code changes:"
    type: text
  - name: footer
    group: page3
    label: "Footer. Information about Breaking Changes and reference issues that this commit closes:"
    type: text
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .hasbreakingbody}}\n\n{{.}}{{end}}{{with .nobreakingbody}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"`
