package main

import (
	"strconv"

	"github.com/Soulou/acadock-monitoring/cpu"
	"github.com/Soulou/acadock-monitoring/mem"
	"github.com/codegangsta/martini"
)

func containerMemUsageHandler(params martini.Params) (int, string) {
	id := params["id"]

	containerMemory, err := mem.GetUsage(id)
	if err != nil {
		return 500, err.Error()
	}
	containerMemoryStr := strconv.FormatInt(containerMemory, 10)
	return 200, containerMemoryStr
}

func containerCpuUsageHandler(params martini.Params) (int, string) {
	id := params["id"]

	containerCpu, err := cpu.GetUsage(id)
	if err != nil {
		return 200, err.Error()
	}
	containerCpuStr := strconv.FormatInt(containerCpu, 10)
	return 200, containerCpuStr
}

func main() {
	go cpu.Monitor()
	r := martini.Classic()
	r.Get("/containers/:id/mem", containerMemUsageHandler)
	r.Get("/containers/:id/cpu", containerCpuUsageHandler)

	r.Run()
}
