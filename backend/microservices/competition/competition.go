package main

import (
	"flag"
	"fmt"

	"backend/microservices/competition/internal/config"
	"backend/microservices/competition/internal/server"
	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/competition.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		competition.RegisterCompetitionServer(grpcServer, server.NewCompetitionServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting microservices server at %s...\n", c.ListenOn)
	s.Start()
}
