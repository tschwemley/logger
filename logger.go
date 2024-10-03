package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

type LoggerOptions struct {
	Debug   bool
	Exclude []string
	LogFile string
}

func SetupLogger(opts LoggerOptions) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if opts.Debug {
		opts.Exclude = []string{"time", "level"}
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	file := os.Stderr // use stderr by default
	if opts.LogFile != "" {
		file, _ = os.OpenFile(
			opts.LogFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
	}

	// prettifies the output
	log = log.Output(zerolog.ConsoleWriter{
		PartsExclude: opts.Exclude,
		Out:          file,
	})
}

func Debug(a ...any) {
	log.Debug().Msg(fmt.Sprint(a...))
}

func Debugf(str string, a ...any) {
	log.Debug().Msg(fmt.Sprintf(str, a...))
}

func Error(a ...any) {
	log.Error().Msg(fmt.Sprint(a...))
}

func Errorf(str string, a ...any) {
	log.Error().Msg(fmt.Sprintf(str, a...))
}

func Fatal(a ...any) {
	log.Fatal().Msg(fmt.Sprint(a...))
}

func Fatalf(str string, a ...any) {
	log.Fatal().Msg(fmt.Sprintf(str, a...))
}

func Info(a ...any) {
	log.Info().Msg(fmt.Sprint(a...))
}

func Infof(str string, a ...any) {
	log.Info().Msg(fmt.Sprintf(str, a...))
}
