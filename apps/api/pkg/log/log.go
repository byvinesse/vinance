package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
)

func SetOutput(w io.Writer) {
	logger = logger.Output(w)
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	return logger.Panic()
}
