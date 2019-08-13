package log

import (
	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Entry is an abstraction over the usual logger
// it adds two more methods "WithFields" and "SystemErr"
type Entry interface {
	// Basic Logging Operation
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	// Attach contextual information to the logging entry
	WithFields(fields map[string]interface{}) Entry

	// SystemErr is a method that disects the error
	// and logs the appropriate level and fields for it.
	SystemErr(err error)
}

type entry struct {
	*logrus.Entry
}

func (e *entry) WithFields(fields map[string]interface{}) Entry {
	ent := e.Entry.WithFields(fields)
	return &entry{ent}
}

func (e *entry) SystemErr(err error) {
	larissaErr, ok := err.(errors.Error)
	if !ok {
		e.Error(err)
		return
	}

	ent := e.WithFields(errFields(larissaErr))
	switch errors.Severity(err) {
	case logrus.WarnLevel:
		ent.Warnf("%v", err)
	case logrus.InfoLevel:
		ent.Infof("%v", err)
	case logrus.DebugLevel:
		ent.Debugf("%v", err)
	default:
		ent.Errorf("%v", err)
	}
}

func errFields(err errors.Error) logrus.Fields {
	f := logrus.Fields{}
	f["operation"] = err.Op
	f["kind"] = errors.KindText(err)
	f["bucket"] = err.Bucket
	f["object"] = err.Object
	f["ops"] = errors.Ops(err)

	return f
}
