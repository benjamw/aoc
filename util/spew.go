package util

import (
	"github.com/davecgh/go-spew/spew"
)

// Dump DON'T USE THIS. It's just a holder for the spew package
func Dump(T interface{}) {
	spew.Dump(T)
}
