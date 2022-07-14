package log_toolkit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	logger "gitlab.falabella.com/fif/integracion/forthehorde/commons/log-toolkit"
)

func TestOsInformation(t *testing.T) {
	os := logger.NewOsInformation()

	t.Run("Valida que el hostname no sea Nulo y que sea de tipo string", func(t *testing.T) {
		var hostname string
		assert.NotNil(t, os.Hostname)
		assert.IsType(t, hostname, os.Hostname)
	})

	t.Run("Valida que el hostname no sea Nulo y que sea de tipo Int", func(t *testing.T) {
		var pid int
		assert.NotNil(t, os.Pid)
		assert.IsType(t, pid, os.Pid)
	})
}
