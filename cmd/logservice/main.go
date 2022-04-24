package main

import (
	"context"
	"fmt"
	"github.com/spjiang/go-test-distributed/log"
	"github.com/spjiang/go-test-distributed/registry"
	"github.com/spjiang/go-test-distributed/service"
	stlog "log"
)

func main() {
	log.Run("/Users/jiangshengping/wwwroot/spjiang/go-test-distributed/cmd/logservice/distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("shutting done log service.")
}
