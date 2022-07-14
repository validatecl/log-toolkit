package log_toolkit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	logger "gitlab.falabella.com/fif/integracion/forthehorde/commons/log-toolkit"
)

func TestBaseConfigurer(t *testing.T) {
	lcfg := logger.NewBaseConfig()

	t.Run("Default configuration", func(t *testing.T) {
		log := lcfg.GenerateConfig(&logger.ZapConfigInput{})
		assert.NotNil(t, log)

	})

	t.Run("Custom configuration", func(t *testing.T) {
		log := lcfg.GenerateConfig(&logger.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "info",
			ConsoleJSONFormat: true,
			EnableFile:        true,
			FileLevel:         "fatal",
			FileJSONFormat:    true,
			MaxSize:           5,
			Filename:          "prueba",
			MaxBackups:        5,
		})
		assert.NotNil(t, log)

	})

	t.Run("Custom configuration", func(t *testing.T) {
		log := lcfg.GenerateConfig(&logger.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "info",
			ConsoleJSONFormat: true,
			EnableFile:        true,
			FileLevel:         "fatal",
			FileJSONFormat:    true,
		})
		assert.NotNil(t, log)

	})

	t.Run("Custom configuration 2", func(t *testing.T) {
		log := lcfg.GenerateConfig(&logger.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "warn",
			ConsoleJSONFormat: true,
			EnableFile:        true,
			FileLevel:         "debug",
			FileJSONFormat:    true,
		})
		assert.NotNil(t, log)

	})

	t.Run("Custom configuration 3", func(t *testing.T) {
		log := lcfg.GenerateConfig(&logger.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "error",
			ConsoleJSONFormat: false,
			EnableFile:        true,
			FileLevel:         "aaaaaa",
			FileJSONFormat:    true,
		})
		assert.NotNil(t, log)
	})

}
