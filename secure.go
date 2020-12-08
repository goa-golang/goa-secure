package goasecure

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

func Md5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	md5Data := hash.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func Hmac(key, data string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func Sha1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func Base64Encode(data string) string {
	sbs := []byte(data)
	encoded := base64.StdEncoding.EncodeToString(sbs)
	return encoded
}

func Base64Decode(data string) (string, error) {
	decodedRaw, err := base64.StdEncoding.DecodeString(data)
	decoded := string(decodedRaw)
	return decoded, err
}
