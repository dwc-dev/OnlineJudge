package main

import (
	"flag"
	"fmt"
	"net/http"

	"user/internal/config"
	"user/internal/handler"
	"user/internal/response"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//封装统一的错误响应
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *response.RespError:
			return e.HttpCode, response.Fail(e)
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
