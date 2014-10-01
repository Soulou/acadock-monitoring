package main

import (
	"flag"
	"log"
	"strconv"

	"net/http/pprof"

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
	doProfile := flag.Bool("profile", false, "profile app")
	flag.Parse()
	go cpu.Monitor()
	r := martini.Classic()

	r.Get("/containers/:id/mem", containerMemUsageHandler)
	r.Get("/containers/:id/cpu", containerCpuUsageHandler)

	if *doProfile {
		log.Println("Enable profiling")
		r.Get("/debug/pprof", pprof.Index)
		r.Get("/debug/pprof/cmdline", pprof.Cmdline)
		r.Get("/debug/pprof/profile", pprof.Profile)
		r.Get("/debug/pprof/symbol", pprof.Symbol)
		r.Post("/debug/pprof/symbol", pprof.Symbol)
		r.Get("/debug/pprof/block", pprof.Handler("block").ServeHTTP)
		r.Get("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
		r.Get("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
		r.Get("/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
	}

	r.Run()
}
