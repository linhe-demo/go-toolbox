package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"runtime"
	"toolbox/exception"
	"toolbox/internal/config"
	"toolbox/internal/handler"
	"toolbox/internal/svc"
)

func main() {
	var configFile *string
	sysType := runtime.GOOS

	if sysType == `linux` {
		configFile = flag.String("l", "/home/go-toolbox/toolbox/etc/toolbox-api.yaml", "the config file")
	} else {
		configFile = flag.String("f", "etc/toolbox-api.yaml", "the config file")
	}
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *exception.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
