package cfg

import "time"

type DataBaseConfig struct {
	DataBaseHost     string
	DataBasePort     string
	DataBaseName     string
	DataBaseUser     string
	DataBasePassword string
	DataBaseProtocol string
}

type AuthConfig struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

type ServerConfig struct {
	ServerHost string
	ServerPort string
}

type Cfg struct {
	ServerCfg   ServerConfig
	DataBaseCfg DataBaseConfig
	AuthCfg     AuthConfig
}

func New(dataBaseCfg DataBaseConfig, serverCfg ServerConfig, authCfg AuthConfig) Cfg {
	newCfg := Cfg{
		ServerCfg:   serverCfg,
		DataBaseCfg: dataBaseCfg,
		AuthCfg:     authCfg,
	}

	return newCfg
}
