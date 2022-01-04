//go:binary-only-package


package sequencemng

import (
	_ "swan/librarys/log"
	_ "swan/librarys/database"
	_ "errors"
	_ "sync"
	_ "time"
	_ "swan/services"
	_ "runtime/debug"
)

func GetSequenceByName(name string)  (seq int,err error){
	return
}

func ResetSequenceByName(name string) (err error) {
	return
}