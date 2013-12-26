package main

import (
	"github.com/Soulou/acadock-live-lxc/lxc/cpu"
	"github.com/Soulou/acadock-live-lxc/lxc/mem"
	"github.com/codegangsta/martini"
	"os"
	"strconv"
)

func containerMemUsageHandler(params martini.Params) (int, string) {
	name := params["name"]

	containerMemory, err := mem.GetUsage(name)
	if err != nil {
		return 500, err.Error()
	}
	containerMemoryStr := strconv.FormatInt(containerMemory, 10)
	return 200, containerMemoryStr
}

func containerCpuUsageHandler(params martini.Params) (int, string) {
	name := params["name"]

	containerCpu, err := cpu.GetUsage(name)
	if err != nil {
		return 200, err.Error()
	}
	containerCpuStr := strconv.FormatInt(containerCpu, 10)
	return 200, containerCpuStr
}

func main() {
	go cpu.Monitor()
	r := martini.Classic()
	r.Get("/containers/:name/mem", containerMemUsageHandler)
	r.Get("/containers/:name/cpu", containerCpuUsageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "4244")
	}
	r.Run()
}
