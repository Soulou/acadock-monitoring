package main

import (
	"Acadock/lxc/cpu"
	"Acadock/lxc/mem"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func containerMemUsageHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Request MEM from : ", req.RemoteAddr)
	vars := mux.Vars(req)
	name := vars["name"]

	containerMemory, err := mem.GetUsage(name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	containerMemoryStr := strconv.FormatInt(containerMemory, 10)
	w.Write([]byte(containerMemoryStr))
}

func containerCpuUsageHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Request CPU from : ", req.RemoteAddr)
	vars := mux.Vars(req)
	name := vars["name"]

	containerCpu := cpu.GetUsage(name)
	containerCpuStr := strconv.FormatInt(containerCpu, 10)
	w.Write([]byte(containerCpuStr))
}

func main() {
	go cpu.Monitor()
	router := mux.NewRouter()
	router.HandleFunc("/containers/{name}/mem", containerMemUsageHandler)
	router.HandleFunc("/containers/{name}/cpu", containerCpuUsageHandler)
	http.ListenAndServe(":4244", router)
}
