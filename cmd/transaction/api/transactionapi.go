package main

import (
	"flag"
	"fmt"
	"runtime"

	"stock/cmd/transaction/api/internal/config"
	"stock/cmd/transaction/api/internal/handler"
	"stock/cmd/transaction/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile *string

func init() {
	var filePath string
	switch runtime.GOOS {
	case "windows":
		filePath = "D:\\Code\\2024\\stock\\cmd\\transaction\\api\\etc\\transactionapi.yaml"
	case "linux":
		filePath = "/etc/stock/transactionapi.yaml"
	}

	configFile = flag.String("f", filePath, "the config file")
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
