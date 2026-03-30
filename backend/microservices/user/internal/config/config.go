package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql   Mysql
	JWT     JWT
	Minio   Minio
	MyRedis MyRedis
}

type MyRedis struct {
	Host     string
	Password string
	DB       int
}

type Mysql struct {
	DataSource string
}

type JWT struct {
	SecretKey string
}

type Minio struct {
	Address   string
	Endpoint  string
	UseSSL    bool
	AccessKey string
	SecretKey string
	Bucket    string
	Avatar    struct {
		Prefix  string
		Default string
	}
}
