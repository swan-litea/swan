//go:binary-only-package

package filecfg

import (
	_ "sync"
	_ "encoding/json"
	_ "os"
	_ "strings"
	_ "io"
	_ "errors"
	_ "sort"
	_ "strconv"
	_ "reflect"
	_ "crypto/md5"
	_ "swan/librarys/log"
	_ "swan/librarys/encryption"
	_ "path"
	_ "encoding/hex"
	_ "path/filepath"
	_ "math/rand"
	_ "time"
	_ "fmt"
	_ "testing"
	_ "swan/librarys/database"
								)

func SetFileName(fullfilename string) (err error) {
	return
}

func Reload() (err error) {
	return nil
}


func ReloadFile(filekey []byte,filename string) (m map[string]interface{},err error) {
	return
}

func Save() (err error) {
	return
}

func LoadServiceCfg(reloadfunc func(map[string]interface{},string) error,servicename string, triggernodes []string) (err error){
	return
}

func UnloadServiceCfg(servicename string) {

}

func SortNodeCfgs(source interface{},sortinfo...string) (rst []map[string]interface{},err error) {
	return
}

func FilterCfgs(source interface{},conditions ...interface{}) (rst []map[string]interface{},err error) {
	return
}

func MapCopy(value interface{}) (newvalue interface{}) {
	return
}

func GetAllCfg() (cfgs map[string]interface{}) {
	return
}