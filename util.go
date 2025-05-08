package utils

import (
	"encoding/json"
)

func ToJSONStr(v any) string {
	if v == nil {
		return ""
	}
	data, _ := json.Marshal(v) //nolint:errchkjson
	return string(data)
}
