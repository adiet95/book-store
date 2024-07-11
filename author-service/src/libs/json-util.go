package libs

import (
	"bytes"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

var (
	jsonIter      = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONMarshal   = json.Marshal
	JSONUnmarshal = json.Unmarshal
)

func mapper() jsoniter.API {
	return jsoniter.ConfigCompatibleWithStandardLibrary
}

func ToJSON(data interface{}) ([]byte, error) {
	return mapper().Marshal(data)
}

func FromJSON(json []byte, result interface{}) error {
	return mapper().Unmarshal(json, result)
}

func Stringify(data interface{}) (string, error) {
	return mapper().MarshalToString(data)
}
func CompactJSON(jsonValue []byte) string {
	requestJSON := new(bytes.Buffer)
	if err := json.Compact(requestJSON, jsonValue); err != nil {
		return "error compact json"
	}
	return requestJSON.String()
}
func ToJSONFromInterface(data map[string]interface{}) ([]byte, error) {
	result, err := json.Marshal(data)
	return result, err
}

func MarshalJSONToString(data interface{}) (string, error) {
	return jsonIter.MarshalToString(data)
}
