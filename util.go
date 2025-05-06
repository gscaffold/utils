package utils

import (
	"encoding/json"
)

func ToJSONStr(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
