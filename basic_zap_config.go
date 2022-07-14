package log_toolkit

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	customTimeFormat string
	pathSeparator    = "/"
)

const (
	defaultLevel      = zapcore.InfoLevel
	defaultTimeFormat = "2006-02-01 15:04:05.000"
	defaultFilePath   = "logs"
	maxSize           = 100

	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

// ZapConfigInput representa a zap config input
type ZapConfigInput struct {
	EnableConsole     bool
	ConsoleLevel      string
	ConsoleJSONFormat bool
	EnableFile        bool
	FilePath          string
	FileLevel         string
	FileJSONFormat    bool
	Filename          string
	MaxSize           int
	MaxBackups        int
	MaxAge            int
	Compress          bool
	TimeFormat        string
}

type zapLogger struct {
	sugaredLogger *zap.Logger
}

// ZapBaseConfigurer generate Zap Config
type ZapBaseConfigurer interface {
	GenerateConfig(input *ZapConfigInput) *zap.Logger
}

type zapBaseConfigurer struct {
}

// NewBaseConfig ...
func NewBaseConfig() ZapBaseConfigurer {
	return &zapBaseConfigurer{}
}

func (z *zapBaseConfigurer) GenerateConfig(input *ZapConfigInput) *zap.Logger {
	return z.newZapCore(input)
}

func (z *zapBaseConfigurer) newZapCore(input *ZapConfigInput) *zap.Logger {

	cores := []zapcore.Core{}
	config := z.validateZapConfig(input)

	if input.EnableConsole {
		level := getZapLevel(input.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(input.ConsoleJSONFormat, config.TimeFormat), writer, level)
		cores = append(cores, core)
	}

	if input.EnableFile {
		level := getZapLevel(input.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.FilePath + pathSeparator + config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		})
		core := zapcore.NewCore(getEncoder(config.FileJSONFormat, config.TimeFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	return zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)
}

func (z *zapBaseConfigurer) validateZapConfig(input *ZapConfigInput) *ZapConfigInput {
	if len(input.Filename) == 0 {
		input.Filename = time.Now().Format("02-01-2006") + ".log"
		log.Printf("default Filename value: " + input.Filename)
	}
	if input.MaxSize <= 0 || input.MaxSize > maxSize {
		input.MaxSize = maxSize
		log.Printf("default MaxSize value: %v MB ", input.MaxSize)
	}
	if len(input.TimeFormat) == 0 {
		input.TimeFormat = defaultTimeFormat
		log.Printf("default TimeFormat value: %s ", input.TimeFormat)
	}
	if len(input.FilePath) == 0 {
		input.FilePath = defaultFilePath
		log.Printf("default FilePath value: %s ", input.FilePath)
	}
	return input
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Debug:
		return zapcore.DebugLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return defaultLevel
	}
}

func getEncoder(isJSON bool, timeFormat string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	customTimeFormat = timeFormat
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.TimeKey = "date"
	encoderConfig.MessageKey = "payload"
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(customTimeFormat))
}
