package service

import (
	"context"
	"fmt"
	"github.com/spjiang/go-test-distributed/registry"
	"log"
	"net/http"
)

/*
 * 服务启动
 *
 */

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	// http router 注册
	registerHandlersFunc()
	// 启动服务
	ctx = startService(ctx, reg.ServiceName, host, port)
	// 服务注册
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	// http server
	var srv http.Server
	srv.Addr = ":" + port
	// 启动HTTP服务
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// 监听终端输入信号
	go func() {
		fmt.Printf("%v started. Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
