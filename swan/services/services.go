//go:binary-only-package

package services

import(
	_ "sync"
	_ "runtime"
)

func NewService(name string,fun interface{}) {
}