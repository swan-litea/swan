//go:binary-only-package

package encryption

import (
	_ "crypto/aes"
	_ "crypto/cipher"
	_ "bytes"
	_ "errors"
	_ "strconv"
	_ "encoding/hex"
	_ "reflect"
	_ "testing"
	_ "fmt"
	_ "golang.org/x/text/encoding/simplifiedchinese"
)

func AesEncrypt(origData, key []byte) (crypted []byte, err error) {
	return
}

func AesDecrypt(crypted, key []byte) (ori []byte, err error) {
	return
}

func ZeroPadding(ciphertext []byte, blockSize int) (rst []byte) {
	return
}

func ZeroUnPadding(origData []byte) (rst []byte) {
	return
}

func PKCS5Padding(ciphertext []byte, blockSize int) (rst []byte) {
	return
}

func PKCS5UnPadding(origData []byte) (rst []byte) {
	return
}

func ENCODE(code string,src interface{}) (rst []byte,err error) {
	return
}

func DECODE(code string,src []byte) (rst interface{},err error) {
	return
}