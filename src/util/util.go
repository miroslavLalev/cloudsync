package util

import (
	"encoding/base64"
)

func EncodePath(path string) string {
	return base64.StdEncoding.EncodeToString([]byte(path))
}

func DecodePath(encodedPath string) string {
	bytes, _ := base64.StdEncoding.DecodeString(encodedPath)
	return string(bytes)
}
