package spewdump

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	facade_consoler "github.com/totoval/consoler/pkg/facade"
	"github.com/totoval/consoler/pkg/structs"
	"github.com/totoval/dumper/internal"
	facade_logger "github.com/totoval/logger/pkg/facade"
	"os"
)

type Dump struct {
	logger   facade_logger.Logger
	consoler facade_consoler.Consoler
}

func (d *Dump) New() internal.Dumper {
	return d
}

func (d *Dump) SetLogger(logger facade_logger.Logger) error {
	d.logger = logger
	return nil
}
func (d *Dump) SetConsoler(consoler facade_consoler.Consoler) error {
	d.consoler = consoler
	return nil
}

func (d *Dump) Dump(v ...interface{}) {
	d.consoler.Println(structs.ColorError, spew.Sdump(v...))
	d.debugPrint(errors.New("====== Totoval Debug ======"))
}

func (d *Dump) DD(v ...interface{}) {
	d.Dump(v...)
	os.Exit(1)
}
