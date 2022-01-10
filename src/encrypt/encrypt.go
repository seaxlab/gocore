package encrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/fnv"
)

// spy 2020/1/21

// Md5 MD5 alg
func Md5(s string) (result string) {
	m := md5.New()
	m.Write([]byte(s))
	b := m.Sum(nil)
	result = hex.EncodeToString(b)
	return
}

// Sha1 SHA1 alg
func Sha1(s string) (result string) {
	sha := sha1.New()
	sha.Write([]byte(s))
	b := sha.Sum(nil)
	result = fmt.Sprintf("%x", b)
	return
}

// Sha256 SHA256 alg
func Sha256(s string) (result string) {
	sha := sha256.New()
	sha.Write([]byte(s))
	b := sha.Sum(nil)
	result = fmt.Sprintf("%x", b)
	return
}

// Sha512 SHA512 alg
func Sha512(s string) (result string) {
	sha := sha512.New()
	sha.Write([]byte(s))
	b := sha.Sum(nil)
	result = fmt.Sprintf("%x", b)
	return
}

// Base64Encode Base64 Encode alg
func Base64Encode(s string) (result string) {
	result = base64.StdEncoding.EncodeToString([]byte(s))
	return
}

// Base64Decode Base64Encode alg
func Base64Decode(s string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "nil", err
	}
	return string(result), nil
}

// FNV32 hashes using fnv32 algorithm
func FNV32(text string) uint32 {
	algorithm := fnv.New32()
	return uint32Hasher(algorithm, text)
}

// FNV32a hashes using fnv32a algorithm
func FNV32a(text string) uint32 {
	algorithm := fnv.New32a()
	return uint32Hasher(algorithm, text)
}

// FNV64 hashes using fnv64 algorithm
func FNV64(text string) uint64 {
	algorithm := fnv.New64()
	return uint64Hasher(algorithm, text)
}

// FNV64a hashes using fnv64a algorithm
func FNV64a(text string) uint64 {
	algorithm := fnv.New64a()
	return uint64Hasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func uint32Hasher(algorithm hash.Hash32, text string) uint32 {
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}

func uint64Hasher(algorithm hash.Hash64, text string) uint64 {
	algorithm.Write([]byte(text))
	return algorithm.Sum64()
}
