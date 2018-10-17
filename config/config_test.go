package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfig(t *testing.T) {
	configuration := NewConfiguration()
	ReadConfig("./mock_config.json", configuration)

	assert.Equal(t, 6, configuration.MiddlewareSize)

}
