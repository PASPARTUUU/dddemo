package meinlogger

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type MeinLogger struct{ *logrus.Logger }

func New() *MeinLogger {
	lg := logrus.New()
	var l MeinLogger
	l.Logger = lg
	return &l
}
func (ml *MeinLogger) Level() log.Lvl {
	lvl := ml.Logger.GetLevel()
	return log.Lvl(lvl)
}
func (ml *MeinLogger) SetLevel(v log.Lvl) {
	ml.Logger.SetLevel(logrus.Level(v))
}

func (ml *MeinLogger) Output() io.Writer {
	return ml.Logger.Out
}
func (ml *MeinLogger) SetOutput(w io.Writer) {
	ml.Logger.SetOutput(w)
}

func (ml *MeinLogger) Prefix() string {
	return "PPrefix"
}
func (ml *MeinLogger) SetPrefix(p string) {}

func (ml *MeinLogger) SetHeader(h string) {}

func (MeinLogger) Debugj(j log.JSON) {}
func (MeinLogger) Infoj(j log.JSON)  {}
func (MeinLogger) Warnj(j log.JSON)  {}
func (MeinLogger) Errorj(j log.JSON) {}
func (MeinLogger) Fatalj(j log.JSON) {}
func (MeinLogger) Panicj(j log.JSON) {}
func (MeinLogger) Printj(j log.JSON) {}
