package spewdump

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/totoval/dumper/internal"
	"github.com/totoval/framework/console"
	"github.com/totoval/logger/pkg/facade"
	"os"
)

type Dump struct {
	logger facade.Logger
}

func (d *Dump) New() internal.Dumper {
	return d
}

func (d *Dump) SetLogger(logger facade.Logger) error {
	d.logger = logger
	return nil
}

func (d *Dump) Dump(v ...interface{}) {
	console.Println(console.CODE_ERROR, spew.Sdump(v...)) //@todo from consoller
	d.debugPrint(errors.New("====== Totoval Debug ======"))
}

func (d *Dump) DD(v ...interface{}) {
	d.Dump(v...)
	os.Exit(1)
}
