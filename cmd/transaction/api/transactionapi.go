package main

import (
	"flag"
	"fmt"

	"stock/cmd/transaction/api/internal/config"
	"stock/cmd/transaction/api/internal/handler"
	"stock/cmd/transaction/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// Linux(代码模板生成的配置)
// var configFile = flag.String("f", "etc/transactionapi.yaml", "the config file")

// Windows(需要修改成绝对路径可执行文件才能正确加载transactionapi.yaml配置文件)
var configFile = flag.String("f", "D:\\YINC_DEVELOPMENT\\go\\stock\\cmd\\transaction\\api\\etc\\transactionapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
