package goasecure

import (
	"fmt"
	"testing"
)

func TestSecure(t *testing.T) {
	fmt.Println(Md5("123"))
	fmt.Println(Hmac("123", "123123"))
	fmt.Println(Base64Decode(Base64Encode("asdfsadf")))
	fmt.Println(Sha1("asdfasdf"))
	s, _ := DesEncrypt("asdfasdf", []byte("asdfasdf"))
	fmt.Println(s)
	fmt.Println(DesDecrypt(s, []byte("asdfasdf")))
}
