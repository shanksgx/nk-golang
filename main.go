package main

import (
	"fmt"
	"nk-golang/routers"
)

func main() {
	// 注册路由
	r := routers.SetupRouter()
	err := r.Run(":8888")
	if err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
		return
	}
}
