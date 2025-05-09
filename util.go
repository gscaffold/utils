package utils

import (
	"encoding/json"
	"log"
)

func ToJSONStr(v any) string {
	if v == nil {
		return ""
	}
	data, _ := json.Marshal(v) //nolint:errchkjson
	return string(data)
}

func HandleFatalError(err error, module string, msg string) {
	if err != nil {
		log.Fatalf("module %s error. err:%s. msg:%s", module, err, msg)
	}
}
