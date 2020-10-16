package radiusd

import (
	"github.com/op/go-logging"
	"github.com/pkg/errors"

	"github.com/ca17/teamsacs/common/log"
)

var radlog = logging.MustGetLogger("Radiusd")

func SetupLog(level logging.Level, syslogaddr string, logdir string, module string) {
	if syslogaddr != "" {
		bf := log.SetupSyslog(level, syslogaddr, module)
		if bf != nil {
			radlog.SetBackend(bf)
		}
	} else {
		bf := log.FileSyslog(level, logdir, module)
		if bf != nil {
			radlog.SetBackend(bf)
		}
	}
}

var (
	Error    = radlog.Error
	Errorf   = radlog.Errorf
	Info     = radlog.Info
	Infof    = radlog.Infof
	Warning  = radlog.Warning
	Warningf = radlog.Warningf
	Fatal    = radlog.Fatal
	Fatalf   = radlog.Fatalf
	Debug    = radlog.Debug
	Debugf   = radlog.Debugf

	IsDebug = func() bool {
		return radlog.IsEnabledFor(logging.DEBUG)
	}
)

func CheckError(err error) {
	if err != nil {
		if IsDebug() {
			panic(errors.WithStack(err))
		} else {
			panic(err)
		}
	}
}
