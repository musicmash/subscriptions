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
  name:  'musicmash-subscriptions'
  login: 'musicmash-subscriptions'
  pass:  'musicmash-subscriptions'
  log: false

log:
  level: DEBUG
  file: 'musicmash-subscriptions.log'

http:
  port: 5566
`
	expected := &AppConfig{
		DB: DBConfig{
			Type:  "mysql",
			Host:  "mariadb",
			Name:  "musicmash-subscriptions",
			Login: "musicmash-subscriptions",
			Pass:  "musicmash-subscriptions",
			Log:   false,
		},
		Log: LogConfig{
			Level:         "DEBUG",
			File:          "musicmash-subscriptions.log",
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
