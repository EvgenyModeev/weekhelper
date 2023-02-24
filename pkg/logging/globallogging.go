package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strings"
	"time"
)

type Logger struct {
	LoggerLevel            string // panic, fatal, error, warn, info, debug, trace
	LoggerMessage          string
	LoggerError            error
	LoggerSubLoggerMessage SubLogger
}

type SubLogger struct {
	SubLoggerLevel   string
	SubLoggerMessage string
	SubLoggerError   error
}

func GlobalLog(generalLog Logger) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Определяем форматы выведенных сообщений.
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	globalLevel, err := zerolog.ParseLevel(generalLog.LoggerLevel)

	globalLogging := zerolog.New(output).With().Timestamp().Logger()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if err != nil {
		// При возникновении ошибки с логгированием выводим ошибку.
		globalLogging.Error().Stack().Err(err).Msg("Возникла проблема с отображением логов.")
	}

	zerolog.SetGlobalLevel(globalLevel)

	switch generalLog.LoggerLevel {
	case "panic":
		globalLogging.Panic().Stack().Err(generalLog.LoggerError).Msg(generalLog.LoggerMessage)
	case "fatal":
		globalLogging.Fatal().Stack().Err(generalLog.LoggerError).Msg(generalLog.LoggerMessage)
	case "error":
		globalLogging.Error().Stack().Err(generalLog.LoggerError).Msg(generalLog.LoggerMessage)
	case "warn":
		globalLogging.Warn().Stack().Err(generalLog.LoggerError).Msg(generalLog.LoggerMessage)
	case "info":
		globalLogging.Info().Stack().Err(generalLog.LoggerError).Msg(generalLog.LoggerMessage)
	case "debug":
		globalLogging.Debug().Stack().Err(generalLog.LoggerError).Msg(generalLog.LoggerMessage)
	case "trace":
		globalLogging.Trace().Stack().Err(generalLog.LoggerError).Msg(generalLog.LoggerMessage)
	default:
		fmt.Printf("не разобрадлся с ошибкой: %s", generalLog.LoggerLevel)
	}

}
