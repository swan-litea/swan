//go:binary-only-package

package base

import (
	_ "encoding/json"
	_ "strconv"
	_ "math/rand"
	_ "time"
		)

func CheckYYYYMMDD24hhmmss(tms string) (valid bool) {
	return
}

func CheckIP(ip string) (valid bool) {
	return
}

func MapCopy(value interface{}) (copyvalue interface{}) {
	return
}

func Realmap(value map[string]interface{}){
}

func Tomap(str string) (rst map[string]interface{},err error) {
	return
}

func Tostringbytes(m map[string]interface{}) (rst []byte) {
	return
}

func RandBytes(len int) (rst []byte) {
	return
}

func RandString(len int) (rst string) {
	return
}

func RandUpperString(len int) (rst string) {
	return
}

func RandLowerString(len int) (rst string) {
	return
}

func RandUpperStringWithNumber(len int) (rst string) {
	return
}

func RandLowerStringWithNumber(len int) (rst string) {
	return
}

func GetSystemTms() string {
	return ""
}

func NextDay(date string) string {
	return ""
}

func PreDay(date string) string {
	return ""
}