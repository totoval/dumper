package spewdump

import (
	"github.com/ztrue/tracerr"
)

func (d *Dump) debugPrint(err error) {
	startFrom := 2
	traceErr := tracerr.Wrap(err)
	frameList := tracerr.StackTrace(traceErr)
	if startFrom > len(frameList) || len(frameList)-2 <= 0 {
		_ = d.logger.Error(err)
	}
	traceErr = tracerr.CustomError(err, frameList[startFrom:len(frameList)-2])
	tracerr.PrintSourceColor(traceErr, 0)
}
