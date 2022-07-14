package log_toolkit_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	log_toolkit "github.com/validatecl/log-toolkit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogToolkit(t *testing.T) {

	log_info := "Prueba Testing Info: Hola Mundo"
	log_debug := "Prueba Testing Debug: Hola Mundo"
	log_warn := "Prueba Testing Warn: Hola Mundo"
	log_error := "Prueba Testing Error: Hola Mundo"
	log_panic := "Prueba Testing Panic: Hola Mundo"

	t.Run("NewLogToolkit creation", func(t *testing.T) {
		core, _ := observer.New(zapcore.InfoLevel)
		logger := zap.New(core)

		logg := log_toolkit.NewLogToolkit(logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "info",
			ConsoleJSONFormat: true,
		})

		assert.IsType(t, &log_toolkit.Logger{}, logg)

	})

	t.Run("Info Log default logger", func(t *testing.T) {
		_, recorded := observer.New(zapcore.InfoLevel)

		logg := log_toolkit.NewLogToolkit(nil, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "info",
			ConsoleJSONFormat: true,
		})

		logg.Info(log_info)
		logg.Debug(log_debug)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {

			assert.NotEqual(t, zapcore.DebugLevel, logs.Level)

			switch logs.Level {
			case zapcore.InfoLevel:
				assert.Equal(t, log_info, logs.Message)
			case zapcore.ErrorLevel:
				assert.Equal(t, log_error, logs.Message)
			case zapcore.WarnLevel:
				assert.Equal(t, log_warn, logs.Message)
			}
		}
	})
	t.Run("Debug Log default logger", func(t *testing.T) {
		_, recorded := observer.New(zapcore.DebugLevel)
		logg := log_toolkit.NewLogToolkit(nil, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "debug",
			ConsoleJSONFormat: true,
		})

		logg.Debug(log_debug)
		logg.Info(log_info)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {
			switch logs.Level {
			case zapcore.DebugLevel:
				assert.Equal(t, log_debug, logs.Message)
			case zapcore.InfoLevel:
				assert.Equal(t, log_info, logs.Message)
			case zapcore.ErrorLevel:
				assert.Equal(t, log_error, logs.Message)
			case zapcore.WarnLevel:
				assert.Equal(t, log_warn, logs.Message)
			}
		}
	})
	t.Run("Warn Log default logger", func(t *testing.T) {
		_, recorded := observer.New(zapcore.WarnLevel)
		logg := log_toolkit.NewLogToolkit(nil, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "warn",
			ConsoleJSONFormat: true,
		})

		logg.Info(log_info)
		logg.Debug(log_debug)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {

			assert.NotEqual(t, zapcore.InfoLevel, logs.Level)
			assert.NotEqual(t, zapcore.DebugLevel, logs.Level)

			switch logs.Level {
			case zapcore.DebugLevel:
			case zapcore.ErrorLevel:
				assert.Equal(t, log_error, logs.Message)
			case zapcore.WarnLevel:
				assert.Equal(t, log_warn, logs.Message)
			}
		}
	})
	t.Run("Error Log default logger", func(t *testing.T) {
		_, recorded := observer.New(zapcore.ErrorLevel)
		logg := log_toolkit.NewLogToolkit(nil, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "error",
			ConsoleJSONFormat: true,
		})

		logg.Info(log_info)
		logg.Debug(log_debug)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {

			assert.NotEqual(t, zapcore.InfoLevel, logs.Level)
			assert.NotEqual(t, zapcore.DebugLevel, logs.Level)
			assert.NotEqual(t, zapcore.WarnLevel, logs.Level)
			assert.Equal(t, log_error, logs.Message)

		}
	})
	t.Run("Debug Log", func(t *testing.T) {
		core, recorded := observer.New(zapcore.DebugLevel)
		logger := zap.New(core)
		logg := log_toolkit.NewLogToolkit(logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "debug",
			ConsoleJSONFormat: true,
		})

		logg.Debug(log_debug)
		logg.Info(log_info)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {
			switch logs.Level {
			case zapcore.DebugLevel:
				assert.Equal(t, log_debug, logs.Message)
			case zapcore.InfoLevel:
				assert.Equal(t, log_info, logs.Message)
			case zapcore.ErrorLevel:
				assert.Equal(t, log_error, logs.Message)
			case zapcore.WarnLevel:
				assert.Equal(t, log_warn, logs.Message)
			}
		}
	})
	t.Run("Info Log", func(t *testing.T) {
		core, recorded := observer.New(zapcore.InfoLevel)
		logger := zap.New(core)

		logg := log_toolkit.NewLogToolkit(logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "info",
			ConsoleJSONFormat: true,
		})

		logg.Info(log_info)
		logg.Debug(log_debug)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {

			assert.NotEqual(t, zapcore.DebugLevel, logs.Level)

			switch logs.Level {
			case zapcore.InfoLevel:
				assert.Equal(t, log_info, logs.Message)
			case zapcore.ErrorLevel:
				assert.Equal(t, log_error, logs.Message)
			case zapcore.WarnLevel:
				assert.Equal(t, log_warn, logs.Message)
			}
		}
	})

	t.Run("Warn Log", func(t *testing.T) {
		core, recorded := observer.New(zapcore.WarnLevel)
		logger := zap.New(core)
		logg := log_toolkit.NewLogToolkit(logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "warn",
			ConsoleJSONFormat: true,
		})

		logg.Info(log_info)
		logg.Debug(log_debug)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {

			assert.NotEqual(t, zapcore.InfoLevel, logs.Level)
			assert.NotEqual(t, zapcore.DebugLevel, logs.Level)

			switch logs.Level {
			case zapcore.DebugLevel:
			case zapcore.ErrorLevel:
				assert.Equal(t, log_error, logs.Message)
			case zapcore.WarnLevel:
				assert.Equal(t, log_warn, logs.Message)
			}
		}
	})
	t.Run("Error Log", func(t *testing.T) {
		core, recorded := observer.New(zapcore.ErrorLevel)
		logger := zap.New(core)
		logg := log_toolkit.NewLogToolkit(logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "error",
			ConsoleJSONFormat: true,
		})

		logg.Info(log_info)
		logg.Debug(log_debug)
		logg.Warn(log_warn)
		logg.Error(log_error)

		for _, logs := range recorded.All() {
			assert.NotEqual(t, zapcore.InfoLevel, logs.Level)
			assert.NotEqual(t, zapcore.DebugLevel, logs.Level)
			assert.NotEqual(t, zapcore.WarnLevel, logs.Level)
			assert.Equal(t, log_error, logs.Message)
		}
	})
	t.Run("Panic Log", func(t *testing.T) {

		core, recorded := observer.New(zapcore.FatalLevel)
		logger := zap.New(core)

		logg := log_toolkit.NewLogToolkit(logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "panic",
			ConsoleJSONFormat: true,
		})

		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				fmt.Printf("%v", err)
				for _, logs := range recorded.All() {

					assert.NotEqual(t, zapcore.InfoLevel, logs.Level)
					assert.NotEqual(t, zapcore.DebugLevel, logs.Level)
					assert.NotEqual(t, zapcore.WarnLevel, logs.Level)
					assert.NotEqual(t, zapcore.ErrorLevel, logs.Level)
					assert.Equal(t, zapcore.PanicLevel, logs.Level)
				}
			}
		}()

		logg.Info(log_info)
		logg.Debug(log_debug)
		logg.Warn(log_warn)
		logg.Error(log_error)
		logg.Panic(log_panic)
	})

}

func TestLogToolkitWithField(t *testing.T) {

	t.Run("Prueba de Logs con Fields por defecto", func(t *testing.T) {
		core, recorded := observer.New(zapcore.InfoLevel)
		logger := zap.New(core)
		info := "Prueba Testing Fields: Hola Mundo"
		expect := "Prueba Testing Fields: Hola Mundo"
		log := log_toolkit.NewLogToolkit(logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "info",
			ConsoleJSONFormat: true,
		})
		initValues := log_toolkit.Fields{
			"channelId": "WEB",
			"country":   "CL",
		}
		logFields := log.WithFields(initValues)

		logFields.Info(info)

		for _, logs := range recorded.All() {
			fmt.Println(logs.Message)
			assert.Equal(t, expect, logs.Message)
			assert.Equal(t, "channelId", logs.Context[0].Key)
			assert.Equal(t, "WEB", logs.Context[0].String)
			assert.Equal(t, "country", logs.Context[1].Key)
			assert.Equal(t, "CL", logs.Context[1].String)

		}

	})
}
