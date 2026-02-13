package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(useJSON bool) {
	if useJSON {
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
		return
	}
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}

func L() *zerolog.Logger {
	l := log.Logger
	return &l
}

func Output(w io.Writer) zerolog.Logger {
	return zerolog.New(w).With().Timestamp().Logger()
}
