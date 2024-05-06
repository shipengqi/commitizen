package parameter

import (
	"regexp"
	"strings"

	"github.com/shipengqi/golib/strutil"
)

var (
	regexName = regexp.MustCompile(`^[a-zA-Z0-9-_]{1,62}$`)
)

const (
	TypeBoolean   = "boolean"
	TypeInteger   = "integer"
	TypeList      = "list"
	TypeMultiList = "multi_list"
	TypeString    = "string"
	TypeSecret    = "secret"
	TypeText      = "text"
)

const UnknownGroup = "unknown"

type FieldKey string

func NewFiledKey(item string, group ...string) FieldKey {
	if len(group) < 1 {
		return FieldKey(UnknownGroup + "." + item)
	}
	if strutil.IsEmpty(group[0]) {
		return FieldKey(UnknownGroup + "." + item)
	}
	return FieldKey(group[0] + "." + item)
}

func GetFiledKey(key string) FieldKey {
	keys := strings.Split(key, ".")
	if len(keys) < 1 {
		return ""
	}
	if len(keys) < 2 {
		return FieldKey(UnknownGroup + "." + keys[0])
	}
	return FieldKey(key)
}
