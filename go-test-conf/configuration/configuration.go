package configuration

import "gopkg.in/gcfg.v1"

type Configuration struct {
	Server struct {
		Port     string
		Protocol string
	}
	Logs struct {
		ErrorLogsPath  string
		AccessLogsPath string
	}
	Database struct {
		ConnectionString   string
		MaxIdleConnections int
	}
}

const defaultConfiguration = "configuration/config.gcfg"

func NewConfiguration(path string) (*Configuration, error) {
	if "" == path {
		path = defaultConfiguration
	}

	c := &Configuration{}
	err := gcfg.ReadFileInto(c, path)
	return c, err
}
