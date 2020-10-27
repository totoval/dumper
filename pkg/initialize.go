package pkg

import (
	"errors"
	"github.com/totoval/dumper/internal"

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
	logger, ok := configuration["logger"].(facade.Logger)
	if !ok {
		return errors.New("unknown configuration logger")
	}

	return d.dump.New().SetLogger(logger)
}
