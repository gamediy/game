package main

import (
	"fmt"
	"game/app/gateway/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
	"os"
	"runtime/pprof"
)

func main() {
	cpuProfile, err := os.Create("./pprof/cpu_profile")
	if err != nil {
		fmt.Printf("创建文件失败:%s", err.Error())
		return
	}
	defer cpuProfile.Close()

	memProfile, err := os.Create("./pprof/mem_profile")
	if err != nil {
		fmt.Printf("创建文件失败:%s", err.Error())
		return
	}
	defer memProfile.Close()
	//采集CPU信息
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	//采集内存信息
	pprof.WriteHeapProfile(memProfile)

	for i := 0; i < 100; i++ {
		fmt.Println("pprof 工具型测试")
	}
	cmd.Main.Run(gctx.GetInitCtx())

}
