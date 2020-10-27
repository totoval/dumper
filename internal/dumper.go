package internal

import "github.com/totoval/logger/pkg/facade"

type Dumper interface {
	New() Dumper
	SetLogger(logger facade.Logger) error
	Dump(v ...interface{})
	DD(v ...interface{})
}
