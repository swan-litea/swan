//go:binary-only-package

package database

import (
	_ "container/list"
	_ "database/sql"
	_ "fmt"
	_ "strings"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-oci8"
	_ "errors"
	_ "sync"
	_ "time"
	_ "strconv"
	"database/sql"
	)

type Database struct {
	DBType						string							//数据库类型
	DBIP						string							//数据库地址
	DBPort						int								//数据库端口
	DBName						string							//数据库名称
	DBUserName					string							//数据库用户名称
	DBUserPswd					string							//数据库用户密码
}

const (
	MYSQL = "mysql"
	ORACLE = "oracle"
	DB2 = "db2"
)

func DBTYPE()  string{
	return ""
}

func User() string {
	return ""
}

func SetFaulttoleranttime(tolerant int)  {
}

func SetSocketTimeout(second int)  {
}

func SetMaxOpenConns(num int)  {
}

func SetMaxConcurrency(num int)  {
}

func SetMaxIdleConns(num int)  {
}

func SetDbInfo(user string,pswd string,ip string,port int,name string,db_type string)  {
}

func Connect() (err error) {
	return
}

func Close() (err error) {
	return
}


func Query(query string,param ...interface{}) (records []map[string]interface{},err error) {
	return
}

func QueryOne(query string,param ...interface{}) (record map[string]interface{},err error) {
	return
}

func Execute(exec string,param ...interface{}) (num int64,err error) {
	return
}

func Begin() (tx *sql.Tx,err error){
	return
}

func Commit(tx *sql.Tx) (err error) {
	return
}

func Rollback(tx *sql.Tx) (err error) {
	return
}

func ExecuteTrans(tx *sql.Tx,exec string,param ...interface{}) (num int64,err error) {
	return
}

func QueryOneTrans(tx *sql.Tx,query string,param ...interface{}) (record map[string]interface{},err error) {
	return

}

func QueryTrans(tx *sql.Tx,query string,param ...interface{}) (records []map[string]interface{},err error) {
	return
}

func GetConcurrencymax() (max int) {
	return
}

func GetConcurrencycur() (cur int) {
	return
}

func TOINT(interface{}) (value int) {
	return
}

func TOSTRING(interface{}) (value string) {
	return
}
