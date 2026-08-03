package cryptor

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(_data []byte) string {
	h := md5.New()
	h.Write(_data) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
