package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// Sign OKX API 签名
func Sign(secret, message string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func StringPtr(s string) *string { return &s }

func Ptr[T any](v T) *T { return &v }
