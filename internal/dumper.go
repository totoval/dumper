package internal

import (
	facade_consoler "github.com/totoval/consoler/pkg/facade"
	facade_logger "github.com/totoval/logger/pkg/facade"
)

type Dumper interface {
	New() Dumper
	SetLogger(logger facade_logger.Logger) error
	SetConsoler(consoler facade_consoler.Consoler) error
	Dump(v ...interface{})
	DD(v ...interface{})
}
