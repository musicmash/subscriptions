package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Load(t *testing.T) {
	// arrange
	data := `
---
db:
  type:  'mysql'
  host:  'mariadb'
  name:  'subscriptions'
  login: 'subscriptions'
  pass:  'subscriptions'
  log: false

log:
  level: DEBUG
  file: 'subscriptions.log'

http:
  port: 5566
`
	expected := &AppConfig{
		DB: DBConfig{
			Type:  "mysql",
			Host:  "mariadb",
			Name:  "subscriptions",
			Login: "subscriptions",
			Pass:  "subscriptions",
			Log:   false,
		},
		Log: LogConfig{
			Level:         "DEBUG",
			File:          "subscriptions.log",
		},
		HTTP: HTTPConfig{
			Port: 5566,
		},
	}

	// action
	err := Load([]byte(data))

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, Config)
}
