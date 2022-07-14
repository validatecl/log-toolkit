package log_toolkit

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Fields ...
type Fields map[string]string

//Toolkit ...
type Toolkit interface {
	Operation
	CustomFields
}

//Operation interface
type Operation interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Panic(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

//CustomFields interface
type CustomFields interface {
	WithFields(keyValues Fields) Toolkit
}

//Logger ...
type Logger struct {
	logger     *zap.Logger
	os         *OsInformation
	marshaller FieldMarshaller
}

//NewLogToolkit metodo por el cual se inicializa log toolkit
func NewLogToolkit(logger *zap.Logger, z *ZapConfigInput) Toolkit {
	if logger == nil {
		logger = NewBaseConfig().GenerateConfig(z)
	}
	return &Logger{
		logger:     logger,
		os:         NewOsInformation(),
		marshaller: NewFieldMarshaller(),
	}
}

//Debug function
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, l.marshaller.MarshalFields(args...)...)
}

//Info function
func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, l.marshaller.MarshalFields(args...)...)
}

//Warn function
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, l.marshaller.MarshalFields(args...)...)
}

//Error function
func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, l.marshaller.MarshalFields(args...)...)
}

//Panic function
func (l *Logger) Panic(msg string, args ...interface{}) {
	l.logger.Panic(msg, l.marshaller.MarshalFields(args...)...)
}

//Fatal function
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatal(msg, l.marshaller.MarshalFields(args...)...)
}

func (log *Logger) WithFields(fields Fields) Toolkit {
	var f = make([]zapcore.Field, 0)
	for key, value := range fields {
		f = append(f, zap.Field{
			Key:       key,
			Interface: value,
			Type:      zapcore.StringType,
			String:    fmt.Sprintf("%v", value),
		})
	}
	return &Logger{
		logger:     log.logger.With(f...),
		os:         log.os,
		marshaller: log.marshaller,
	}
}
