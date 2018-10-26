package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

// HashPassword Hash密码
func HashPassword(password, salt string) string {
	ctx := sha256.New()
	ctx.Write([]byte(password))
	hashed := ctx.Sum(nil)
	ctx.Reset()
	ctx.Write(hashed)
	ctx.Write([]byte(salt))
	return hex.EncodeToString(ctx.Sum(nil))
}

// GetRandomSalt 生成随机盐值
func GetRandomSalt() string {
	return getRandomString(16)
}

// GetRandomString 生成随机字符串
func getRandomString(length int) string{
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}