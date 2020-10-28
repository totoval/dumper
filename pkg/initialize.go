package pkg

import (
	"errors"
	"github.com/totoval/dumper/internal"

	facade_consoler "github.com/totoval/consoler/pkg/facade"
	"github.com/totoval/dumper/internal/spewdump"
	"github.com/totoval/logger/pkg/facade"
)

type Dump struct {
	dump internal.Dumper
}

func (d *Dump) Component() interface{} {
	return d.dump
}

func (d *Dump) Use(driver string) Componentor {
	switch driver {
	case DriverSpewDump:
		d.dump = &spewdump.Dump{}
	default:
		d.dump = &spewdump.Dump{}
	}
	return d
}

func (d *Dump) Config(configuration map[string]interface{}) error {
	d.dump.New()

	logger, ok := configuration["logger"].(facade.Logger)
	if !ok {
		return errors.New("unknown configuration logger")
	}
	if err:=d.dump.SetLogger(logger);err != nil{
		return err
	}

	consoler, ok := configuration["consoler"].(facade_consoler.Consoler)
	if !ok {
		return errors.New("unknown configuration consoler")
	}

	if err:=d.dump.SetConsoler(consoler);err != nil{
		return err
	}

	return nil
}
