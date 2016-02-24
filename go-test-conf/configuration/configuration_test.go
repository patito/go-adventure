package configuration_test

import (
	"testing"

	"github.com/JSainsburyPLC/api-notifications/configuration"
	"github.com/stretchr/testify/require"
)

func TestNewConfiguration(t *testing.T) {
	require := require.New(t)
	c, err := configuration.NewConfiguration("config_test.gcfg")
	require.Nil(err)
	require.Equal(":1111", c.Server.Port, "Possibly wrong config or no port set")
	require.Equal("tcp", c.Server.Protocol, "Possibly wrong config or no  set")
	require.Equal("../log_test/api_notifications_error.log",
		c.Logs.ErrorLogsPath, "Possibly wrong config or no error log path set")
	require.Equal("../log_test/api_notifications_access.log",
		c.Logs.AccessLogsPath, "Possibly wrong config or no access log path set")
	require.Equal("postgres://localhost/notifications?sslmode=disable",
		c.Database.ConnectionString, "Possibly wrong config or no database configuration")
	require.Equal(2000,
		c.Database.MaxIdleConnections, "Possibly wrong config or no maxIdleConns set")
}

func TestNewConfigurationDefaultConfig(t *testing.T) {
	require := require.New(t)
	c, err := configuration.NewConfiguration("")
	require.Nil(err)
	require.Equal(":7000", c.Server.Port, "Possibly wrong config or no port set")
	require.Equal("tcp", c.Server.Protocol, "Possibly wrong config or no  set")
	require.Equal("log_test/api_notifications_error.log",
		c.Logs.ErrorLogsPath, "Possibly wrong config or no error log path set")
	require.Equal("log_test/api_notifications_access.log",
		c.Logs.AccessLogsPath, "Possibly wrong config or no access log path set")
	require.Equal("postgres://localhost/notifications?sslmode=disable",
		c.Database.ConnectionString, "Possibly wrong config or no database configuration")
	require.Equal(300,
		c.Database.MaxIdleConnections, "Possibly wrong config or no maxIdleConns set")
}
