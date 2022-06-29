package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySql struct {
		DataSource string
	}
	Redis struct {
		Addr     string
		Password string
	}
}
