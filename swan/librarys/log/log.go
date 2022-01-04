//go:binary-only-package

package log

import (
	_ "os"
	_ "sync"
	_ "runtime"
	_ "bytes"
	_ "strconv"
	_ "time"
	_ "strings"
	_ "fmt"
	_ "errors"
	_ "path"
	_ "log"
)

func SetLogLevel(level string){
}

func SetPath(dirpath string){
}

func NewLog(name string) (err error) {
	return
}

func DestroyLog() {
}

func Errorln(args ...interface{}){
}

func Warnln(args ...interface{}) {
}

func Infoln(args ...interface{}) {
}

func Debugln(args ...interface{}) {
}

func ErrorlnWithDeep(deep int,args ...interface{}) {
}

func WarnlnWithDeep(deep int,args ...interface{}) {
}

func InfolnWithDeep(deep int,args ...interface{}) {
}

func DebuglnWithDeep(deep int,args ...interface{}) {
}

func ErrorlnToFile(logname string,args ...interface{})  {
}

func WarnlnToFile(logname string,args ...interface{})  {
}

func InfolnToFile(logname string,args ...interface{})  {
}

func DebuglnToFile(logname string,args ...interface{})  {
}

func ErrorlnToFileWithDeep(logname string,deep int,args ...interface{})  {
}

func WarnlnToFileWithDeep(logname string,deep int,args ...interface{})  {
}

func InfolnToFileWithDeep(logname string,deep int,args ...interface{})  {
}

func DebuglnToFileWithDeep(logname string,deep int,args ...interface{})  {
}

func GetCurrentLogName() (logname string) {
	return
}