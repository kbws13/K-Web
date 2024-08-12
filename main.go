package main

import (
	"KWeb/app/console"
	"KWeb/app/http"
	"KWeb/framework"
	"KWeb/framework/provider/app"
	"KWeb/framework/provider/distributed"
	"KWeb/framework/provider/kernel"
)

func main() {
	// 初始化容器服务
	container := framework.NewKContainer()
	// 绑定 APP 服务提供者
	container.Bind(&app.KAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&distributed.LocalDistributedProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.KKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
