package main

import (
	"context"
	"fmt"
	"github.com/spjiang/go-test-distributed/registry"
	"log"
	"net/http"
)

func main() {
	// 注册http route , 与http server 绑定 registry.RegistryService
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort
	go func() {
		// 启动HTTP服务
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// 监听终端输入信号
	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()

	fmt.Println("shutting done registry  service.")
}
