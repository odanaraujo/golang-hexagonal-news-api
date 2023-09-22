package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log         *zap.Logger
	LOG_OUTPUT  = "LOG_OUTPUT"
	LOG_LEVEL   = "LOG_LEVEL"
	LEVEL_KEY   = "LEVEL"
	TIME_KEY    = "time"
	MESSAGE_KEY = "message"
	ENCODING    = "json"
)

func init() {
	logconfig := zap.Config{
		OutputPaths: []string{getLogOutput()},            //onde jogar esse log? stdout, por exemplo, joga no terminal
		Level:       zap.NewAtomicLevelAt(getLogLevel()), //qual o nível de log? info, debug, error, etc
		Encoding:    ENCODING,                            //qual o formato do retorno do log? json, console, etc
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     LEVEL_KEY,                          //seja level dentro do json
			TimeKey:      TIME_KEY,                           //seja time dentro do json
			MessageKey:   MESSAGE_KEY,                        //seja message dentro do json
			EncodeTime:   zapcore.ISO8601TimeEncoder,         //seja o formato da data no padrão ISO8601
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder, //padroniza os logs em caixa baixa
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logconfig.Build()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

func getLogOutput() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))

	if output == "" {
		return "stdout"
	}
	return output
}

func getLogLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
