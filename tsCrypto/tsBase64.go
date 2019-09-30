package tsCrypto

import (
	"encoding/base64"
)

func Base64EncodeByte(str []byte) string {
	data := base64.StdEncoding.EncodeToString(str)
	return data
}

func Base64DecodeByte(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return []byte{}
	}
	return data
}
func Base64Encode(str string) string {
	data := base64.StdEncoding.EncodeToString([]byte(str))
	return data
}

func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}
