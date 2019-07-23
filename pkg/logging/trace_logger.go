package logging

import (
	"github.com/sirupsen/logrus"
)

type TraceLogger interface {
	WithField(key string, value interface{}) TraceLogger
	Debug(args ...interface{})
}

/* stub logger if tracing is disabled */
type emptyTraceLogger struct {
}

func (l *emptyTraceLogger) WithField(key string, value interface{}) TraceLogger {
	return l
}

func (l *emptyTraceLogger) Debug(args ...interface{}) {

}

/* trace logger based on logger instance */
type logTraceLogger struct {
	log *logrus.Entry
}

func (l *logTraceLogger) WithField(key string, value interface{}) TraceLogger {
	newLog := l.log.WithField(key, value)
	return &logTraceLogger{
		log: newLog,
	}
}

func (l *logTraceLogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

/* trace logger based on logrus */
type logrusTraceLogger struct{}

func (l *logrusTraceLogger) WithField(key string, value interface{}) TraceLogger {
	newLog := logrus.WithField(key, value)
	return &logTraceLogger{
		log: newLog,
	}
}

func (l *logrusTraceLogger) Debug(args ...interface{}) {
	logrus.WithField("trace", true).Debug(args...)
}

var traceEnabled = false

func GetTraceLogger(log *logrus.Entry) TraceLogger {
	if !traceEnabled {
		return &emptyTraceLogger{}
	}
	return &logTraceLogger{
		log: log.WithField("trace", true),
	}
}

func GetEnabledTraceLogger(log *logrus.Entry) TraceLogger {
	return &logTraceLogger{
		log: log.WithField("trace", true),
	}
}

func NewTraceLogger() TraceLogger {
	if !traceEnabled {
		return &emptyTraceLogger{}
	}
	return &logrusTraceLogger{}
}
