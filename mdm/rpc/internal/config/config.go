package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	PostgresConf struct {
		Host            string
		Port            int
		User            string
		DbName          string
		Password        string
		SslMode         string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifeTime int
		ConnMaxIdleTime int
	}
}
