//go:binary-only-package

package unpack

import (
	_ "swan/librarys/database"
	_ "swan/librarys/encryption"
	_ "swan/librarys/log"
	_ "errors"
	_ "strconv"
	_ "strings"
	_ "fmt"
	_ "sync"
	_ "os"
)

func InitFromDB() (err error) {
	return
}

func InitFromFile() (err error) {
	return
}

func Unpack(msgcode string,src []byte) (rst map[string]interface{},err error) {
	return
}

func UnpackFile(fullfilename string,code string,ignorestart int,ignoreend int) (rst map[string]interface{},err error) {
	return
}

