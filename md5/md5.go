package md5

import "crypto/md5"

// md5 16byte哈希
func CalculateHash(data []byte) [16]byte {
	return md5.Sum(data)
}
