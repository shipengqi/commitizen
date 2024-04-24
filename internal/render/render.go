package render

import (
	"strings"
)

const (
	TypeSelect   = "select"
	TypeInput    = "input"
	TypeTextArea = "textarea"
)

// -----------------------------------------

func isEmptyStr(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}
