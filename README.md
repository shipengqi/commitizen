# commitizen

[![test](https://github.com/shipengqi/commitizen/actions/workflows/e2e.yaml/badge.svg)](https://github.com/shipengqi/commitizen/actions/workflows/e2e.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/shipengqi/commitizen)](https://goreportcard.com/report/github.com/shipengqi/commitizen)
[![release](https://img.shields.io/github/release/shipengqi/commitizen.svg)](https://github.com/shipengqi/commitizen/releases)
[![license](https://img.shields.io/github/license/shipengqi/commitizen)](https://github.com/shipengqi/commitizen/blob/main/LICENSE)

Command line utility to standardize git commit messages, golang version. Forked from [commitizen-go](https://github.com/lintingzhen/commitizen-go).

Fixes some issues of commitizen-go and supports more new features.

![demo](https://github.com/shipengqi/illustrations/blob/e0d588dd70551344f0394cbf6671b15ae22e7635/commitizen/demo.gif?raw=true)

## Features

- Multi-template support.
- More powerful and flexible template.
- Support more options of `git commit`.
- Use [huh](https://github.com/charmbracelet/huh) instead of [survey](https://github.com/AlecAivazis/survey) ([survey](https://github.com/AlecAivazis/survey) is no longer maintained).

## Getting Started

```
Command line utility to standardize git commit messages.

Usage:
  commitizen

Available Commands:
  init        Install this tool to git-core as git-cz.
  version     Print the CLI version information.      
  help        Help about any command

Git Commit flags:
  -a, --all
                commit all changed files.
      --amend
                amend previous commit
      --author string
                override author for commit
      --date string
                override date for commit
  -q, --quiet
                suppress summary after successful commit
  -s, --signoff
                add a Signed-off-by trailer.
  -v, --verbose
                show diff in commit message template

Commitizen flags:
  -d, --default
                use the default template, '--default' has a higher priority than '--template'.
      --dry-run
                do not create a commit, but show the message and list of paths
                that are to be committed.
  -t, --template string
                template name to use when multiple templates exist.

Use "commitizen [command] --help" for more information about a command.
```

Commit with commitizen:

```
$ git cz
```

## Installation

### From the Binary Releases

Download the pre-compiled binaries from the [releases page](https://github.com/shipengqi/commitizen/releases) and copy them to the desired location.

Then install this tool to git-core as git-cz:

```
$ commitizen init
```

### Go Install

You must have a working Go environment:

```
$ go install github.com/shipengqi/commitizen@latest
$ commitizen init
```

### From Source

You must have a working Go environment:

```
$ git clone https://github.com/shipengqi/commitizen.git
$ cd commitizen
$ make && ./_output/$(GOOS)/$(GOARCH)/bin/commitizen init
```

## Configuration

You can set configuration file that `.czrc` at repository root or home directory. The configuration file that located in repository root have a priority over the one in home directory. The format is the same as the following:

```yaml
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
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .hasbreakingbody}}\n\n{{.}}{{end}}{{with .nobreakingbody}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
```

### Default

Optional. If true, the template will be used as the default template, note that there can only be one default template.

### Format

Commit message `format`:

```
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
```

### Items

#### Common Item Properties

| Property    | Required | Default Value | Description                                                                                                         |
|:------------|:---------|:--------------|:--------------------------------------------------------------------------------------------------------------------|
| name        | yes      | -             | Unique identifier for the item.                                                                                     |
| label       | yes      | -             | This will be used as the label for the input field in the UI.                                                       |
| type        | yes      | -             | The type of item. Determines which UI widget is shown. See the Item Types section to see all the different options. |
| group       | no       | -             | The name of the group this item belongs to. Separates items into groups (you can think of groups as pages).         |
| description | no       | -             | A short description of the item for user guidance. This will be displayed along with the input field.               |

#### Item Types

- string
- text
- integer
- boolean
- secret
- list
- multi_list

#### string

`string` are single line text parameters.

Properties:

| Property      | Required | Default Value | Description                                                                                                  |
|:--------------|:---------|:--------------|:-------------------------------------------------------------------------------------------------------------|
| required      | no       | `false`       | Whether a string value is required or not.                                                                   |
| fqdn          | no       | `false`       | Add a preset FQDN regex to validate string.                                                                  |
| ip            | no       | `false`       | Add a preset IPv4/IPv6 regex to validate string.                                                             |
| trim          | no       | `false`       | If true, will remove the leading and trailing blank characters before submit.                                |
| default_value | no       | -             | The default value for this item.                                                                             |
| regex         | no       | -             | A regex used to validate the string.                                                                         |
| min_length    | no       | -             | The minimum length of the string. If the value is not required and no value has been given, this is ignored. |
| max_length    | no       | -             | The maximum length of the string.                                                                            |

#### text

Properties:

| Property      | Required | Default Value | Description                                                                                                |
|:--------------|:---------|:--------------|:-----------------------------------------------------------------------------------------------------------|
| required      | no       | `false`       | Whether the text is required or not.                                                                       |
| height        | no       | 5             | The height of the text.                                                                                    |
| default_value | no       | -             | The default value for this item.                                                                           |
| regex         | no       | -             | A regex used to validate the text.                                                                         |
| min_length    | no       | -             | The minimum length of the text. If the value is not required and no value has been given, this is ignored. |
| max_length    | no       | -             | The maximum length of the text.                                                                            |

#### integer

`integer` is a number.

Properties:

| Property      | Required | Default Value | Description                             |
|:--------------|:---------|:--------------|:----------------------------------------|
| required      | no       | `false`       | Whether the integer is required or not. |
| default_value | no       | -             | The default value for this item.        |
| min           | no       | -             | The minimum value allowed.              |
| max           | no       | -             | The maximum value allowed.              |

#### boolean

`boolean` are true or false values.

Properties:

| Property      | Required | Default Value | Description                      |
|:--------------|:---------|:--------------|:---------------------------------|
| default_value | no       | -             | The default value for this item. |

#### secret

`secret` is used for sensitive data that should not be echoed in the UI, for example, passwords.

Properties:

| Property      | Required | Default Value | Description                                                                                                  |
|:--------------|:---------|:--------------|:-------------------------------------------------------------------------------------------------------------|
| required      | no       | `false`       | Whether the secret is required or not.                                                                       |
| trim          | no       | `false`       | If true, will remove the leading and trailing blank characters before submit.                                |
| default_value | no       | -             | The default value for this item.                                                                             |
| regex         | no       | -             | A regex used to validate the secret.                                                                         |
| min_length    | no       | -             | The minimum length of the secret. If the value is not required and no value has been given, this is ignored. |
| max_length    | no       | -             | The maximum length of the secret.                                                                            |

#### list

`list` is predefined lists of values that can be picked by the user.

Properties:

| Property      | Required | Default Value | Description                                                                                           |
|:--------------|:---------|:--------------|:------------------------------------------------------------------------------------------------------|
| required      | no       | `false`       | Whether a string value is required or not.                                                            |
| default_value | no       | -             | The default value for this item.                                                                      |
| options       | yes      | -             | The list of options to choose from.                                                                   |
| height        | no       | -             | The height of the list. If the number of options exceeds the height, the list will become scrollable. |

#### multi_list

Similar to `list`, but with multiple selection.

Properties:

| Property      | Required | Default Value | Description                                                                                           |
|:--------------|:---------|:--------------|:------------------------------------------------------------------------------------------------------|
| required      | no       | `false`       | Whether a string value is required or not.                                                            |
| default_value | no       | -             | A list of default selection values.                                                                   |
| options       | yes      | -             | The list of options to choose from.                                                                   |
| limit         | no       | `false`       | The limit of the multiple selection list.                                                             |
| height        | no       | -             | The height of the list. If the number of options exceeds the height, the list will become scrollable. |

#### list/multi_list Options

Properties:

| Property | Required | Description                      |
|:---------|:---------|:---------------------------------|
| key      | yes      | The message shown in the UI.     |
| value    | yes      | Unique identifier for the value. |

### Groups (Optional)

Group Properties:

| Property   | Required | Default Value | Description                                                                              |
|:-----------|:---------|:--------------|:-----------------------------------------------------------------------------------------|
| name       | yes      | -             | Unique identifier for the property.                                                      |
| depends_on | no       | -             | If this group should only be shown when a specific condition is met on another property. |

DependsOn Properties:

| Property       | Required | Default Value | Description                                                                                  |
|:---------------|:---------|:--------------|:---------------------------------------------------------------------------------------------|
| or_conditions  | no       | `[]`          | The list of conditions in which at least one must be satisfied for the property to be shown. |
| and_conditions | no       | `[]`          | The list of conditions in which all must be satisfied for the property to be shown.          |

DependsOn Conditions:

- ValueEqualsCondition
- ValueNotEqualsCondition
- ValueContainsCondition
- ValueNotContainsCondition
- ValueEmptyCondition

#### ValueEqualsCondition

Properties:

| Property       | Required | Default Value | Description                                                                                     |
|:---------------|:---------|:--------------|:------------------------------------------------------------------------------------------------|
| parameter_name | yes      | -             | The name of the group that the current group is dependent upon. for example, `page2.isbreaking` |
| value_equals   | yes      | -             | The value the target parameter must equal for this condition to be considered true.             |

#### ValueNotEqualsCondition

Properties:

| Property         | Required | Default Value | Description                                                                             |
|:-----------------|:---------|:--------------|:----------------------------------------------------------------------------------------|
| parameter_name   | yes      | -             | The name of the group that the current group is dependent upon.                         |
| value_not_equals | yes      | -             | The value the target parameter must not equal for this condition to be considered true. |

#### ValueContainsCondition

Properties:

| Property       | Required | Default Value | Description                                                                         |
|:---------------|:---------|:--------------|:------------------------------------------------------------------------------------|
| parameter_name | yes      | -             | The name of the group that the current group is dependent upon.                     |
| value_contains | yes      | -             | A value the target parameter must contain for this condition to be considered true. |

#### ValueNotContainsCondition

Properties:

| Property           | Required | Default Value | Description                                                                             |
|:-------------------|:---------|:--------------|:----------------------------------------------------------------------------------------|
| parameter_name     | yes      | -             | The name of the group that the current group is dependent upon.                         |
| value_not_contains | yes      | -             | A value the target parameter must not contain for this condition to be considered true. |

#### ValueEmptyCondition

Properties:

| Property       | Required | Default Value | Description                                                                      |
|:---------------|:---------|:--------------|:---------------------------------------------------------------------------------|
| parameter_name | yes      | -             | The name of the group that the current group is dependent upon.                  |
| value_empty    | yes      | -             | A bool value reflecting whether the expected parameter should be empty or not.   |

### Multiple Templates

You can define multiple templates in the `.czrc` file, separated by `---`：

```yaml
name: angular-template
items:
# ...  
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"`

---

name: my-template
items:
# ...  
format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"`
```

![multiple-templates](https://github.com/shipengqi/illustrations/blob/e0d588dd70551344f0394cbf6671b15ae22e7635/commitizen/multiple-templates.png?raw=true)