package main

import (
	"flag"
	"fmt"
	"starbucks/starbucks/global"
	"starbucks/starbucks/internal/config"
	"starbucks/starbucks/internal/handler"
	"starbucks/starbucks/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./starbucks/etc/starbucks-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	global.Config = &c

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
