package logging

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

type Logger struct {
	logger log.Logger
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = level.NewFilter(logger, level.AllowInfo()) // <--
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
func GetLogger() log.Logger {
	return logger
}