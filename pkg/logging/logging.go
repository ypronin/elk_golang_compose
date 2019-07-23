package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// type alias
type Fields = logrus.Fields

type LoggerProxy interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}

func Init(serviceName string) *logrus.Entry {
	buildVersion = getBuildVersion()
	runtimeVersion = getRuntimeVersion()

	if os.Getenv("TXTLOG") != "" {
		logrus.SetFormatter(new(logrus.TextFormatter))
	} else {
		logrus.SetFormatter(new(logrus.JSONFormatter))
	}

	data := Fields{
		"build":   buildVersion,
		"runtime": runtimeVersion,
		"service": serviceName,
	}

	logrus.AddHook(&ContextHook{Data: data})

	logrus.Printf("Service %s %s, branch %s, build %s", serviceName, gitRevision, gitBranch, buildNumber)
	logrus.New()
	return logrus.NewEntry(logrus.StandardLogger())
}

func InitFormat(log *logrus.Logger) {
	if os.Getenv("TXTLOG") != "" {
		log.Formatter = new(logrus.TextFormatter)
		return
	}

	log.Formatter = new(logrus.JSONFormatter)
}

type ContextHook struct {
	Data Fields
}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") &&
			!strings.Contains(name, "adapter.go") &&
			!strings.Contains(name, "logging.(*logTraceLogger)") &&
			!strings.Contains(name, "logrus.(*adapter)") &&
			!strings.Contains(name, "logrus.(*Entry)") &&
			!strings.Contains(name, "CloseWithLog") {
			file, line := fu.FileLine(pc[i] - 1)

			data := make(map[string]interface{})
			for k, v := range entry.Data {
				data[k] = v
			}

			data["file"] = fmt.Sprintf("%v:%v", path.Base(file), line)
			data["func"] = path.Base(name)

			for k, v := range hook.Data {
				data[k] = v
			}

			entry.Data = data

			break
		}
	}
	return nil
}

// CloseWithLog close the passed Closer, with error logged if accured
func CloseWithLog(l *logrus.Entry, c io.Closer) {
	err := c.Close()
	if err != nil {
		l.Error(err)
	}
}
