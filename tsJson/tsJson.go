package tsJson

import (
	"encoding/json"
)

func ToJson(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(str)
}

func ToString(obj interface{}) (string, error) {
	str, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func FromString(str string, obj interface{}) (error) {
	err := json.Unmarshal([]byte(str), obj)
	return err
}
