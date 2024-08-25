package main

import (
	"flag"
	"fmt"
	"runtime"
	"stock/cmd/market/api/internal/config"
	"stock/cmd/market/api/internal/handler"
	"stock/cmd/market/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile *string

func init() {
	var filePath string
	switch runtime.GOOS {
	case "windows":
		filePath = "D:\\Code\\2024\\stock\\cmd\\market\\api\\etc\\marketapi.yaml"
	case "linux":
		filePath = "/etc/stock/marketapi.yaml"
	}

	configFile = flag.String("f", filePath, "the config file")
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	origin := fmt.Sprintf("%v:%v", c.Host, c.Port)
	var origins []string
	origins = append(origins, origin)
	// 设置允许跨域：https://github.com/zeromicro/go-zero/issues/422
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
