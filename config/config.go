package config

import "gollow/data"

var MySQLConfig = &data.MysqlConfig{}

func Init() {

	MySQLConfig = &data.MysqlConfig{
		Dsn:     "root:password@/surge_engine?parseTime=true",
		MaxIdle: 1000,
		MaxOpen: 10,
	}
}
