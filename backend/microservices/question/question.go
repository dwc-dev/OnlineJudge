package main

import (
	"flag"
	"fmt"

	"backend/microservices/question/internal/config"
	"backend/microservices/question/internal/server"
	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/question.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		question.RegisterQuestionServer(grpcServer, server.NewQuestionServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting microservices server at %s...\n", c.ListenOn)
	s.Start()
}
