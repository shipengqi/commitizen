package render

import (
	"strings"
)

const (
	TypeSelect    = "select"
	TypeInput     = "input"
	TypeTextArea  = "textarea"
	DescMaxLength = 50
)

// -----------------------------------------

func isEmptyStr(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}
