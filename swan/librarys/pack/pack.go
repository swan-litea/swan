//go:binary-only-package

package pack

import (
	_ "swan/librarys/database"
	_ "swan/librarys/encryption"
	_ "swan/librarys/log"
	_ "errors"
	_ "strconv"
	_ "strings"
	_ "fmt"
	_ "sync"
	_ "path"
	_ "os"
	_ "errors"
	_ "reflect"
	"reflect"
)

func InitFromDB() (err error) {
	return
}

func InitFromFile() (err error) {
	return
}

func Pack(msgcode string,src map[string]interface{}) (rst []byte,err error) {
	return
}

func PackByCode(packcode string,src map[string]interface{}) (rst []byte,err error) {
	return
}

func PackFile(fullfilename string,src map[string]interface{},code string,handle reflect.Value) (err error) {
	return
}