package main

import (
	"flag"
	"fmt"

	"backend/code-sandbox/docker"
	"backend/rpc/judge/internal/config"
	"backend/rpc/judge/internal/server"
	"backend/rpc/judge/internal/svc"
	"backend/rpc/judge/pb/judge"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/judge.yaml", "the config file")

func main() {

	docker.InitClient()

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		judge.RegisterJudgeServer(grpcServer, server.NewJudgeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
