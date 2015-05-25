package config

import (
	"code.google.com/p/gcfg"
	"flag"
	"go.iondynamics.net/iDlogger"
)

var Std *Config

type Config struct {
	Http struct {
		CookieSecret  string
		CookieTimeout int
		CookieSecure  bool
		Listen        string
		Fcgi          bool
	}

	PassPad struct {
		Workspace      string
		DevelopmentEnv bool
		OtpIssuer      string
	}

	Logger struct {
		SlackLogUrl string
	}
}

func init() {
	ConfigPath := flag.String("conf", "./config.gcfg", "Specify a file containing the configuration or \"./config.gcfg\" is used")
	flag.Parse()
	Std = &Config{}
	err := gcfg.ReadFileInto(Std, *ConfigPath)

	if err != nil {
		iDlogger.Fatal("config: ", err)
	}
}
