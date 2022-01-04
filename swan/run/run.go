//go:binary-only-package

package run

import (
	_ "encoding/hex"
	_ "errors"
	_ "fmt"
	_ "github.com/fsnotify/fsnotify"
	_ "github.com/go-ini/ini"
	_ "github.com/hprose/hprose-golang/rpc"
	_ "github.com/shirou/gopsutil/cpu"
	_ "github.com/shirou/gopsutil/mem"
	_ "github.com/sparrc/go-ping"
	_ "net"
	_ "os"
	_ "strconv"
	_ "strings"
	_ "swan/global"
	_ "swan/librarys/base"
	_ "swan/librarys/database"
	_ "swan/librarys/encryption"
	_ "swan/librarys/log"
	_ "swan/librarys/pack"
	_ "swan/librarys/unpack"
	_ "swan/services"
	_ "sync"
	_ "syscall"
	_ "testing"
	_ "time"
)

func Run() {
}
