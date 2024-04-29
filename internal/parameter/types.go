package parameter

import "github.com/shipengqi/golib/strutil"

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

func GetFiledKey(item string, group ...string) FieldKey {
	if len(group) < 1 {
		return FieldKey(UnknownGroup + "." + item)
	}
	if strutil.IsEmpty(group[0]) {
		return FieldKey(UnknownGroup + "." + item)
	}
	return FieldKey(group[0] + "." + item)
}
