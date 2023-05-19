package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func RandSeqStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// md5加密
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// otp token
func GetOtpToken(digits int, secret string, period int) string {
	token, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    uint(period),
		Skew:      1,
		Digits:    otp.Digits(digits),
		Algorithm: otp.AlgorithmSHA1,
	})

	if err != nil {
		panic(err)
	}

	return token
}
