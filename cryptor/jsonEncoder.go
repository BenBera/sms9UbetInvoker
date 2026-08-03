package cryptor

import "encoding/json"

func JsonEncode(_jsonObj interface{}) string {
	jsBytes, _ := json.Marshal(_jsonObj)
	return string(jsBytes)
}
