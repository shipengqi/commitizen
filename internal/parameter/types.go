package parameter

import (
	"github.com/shipengqi/golib/strutil"
	"strings"
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
