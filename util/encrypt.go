package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptToMD5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
