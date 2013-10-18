package main

import (
	"fmt"
	"github.com/Soulou/acadock-live-lxc/lxc/cpu"
	"github.com/Soulou/acadock-live-lxc/lxc/mem"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

type containerMemUsageHandler struct{}

func (c *containerMemUsageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

type containerCpuUsageHandler struct{}

func (c *containerCpuUsageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request CPU from : ", req.RemoteAddr)
	vars := mux.Vars(req)
	name := vars["name"]

	containerCpu := cpu.GetUsage(name)
	containerCpuStr := strconv.FormatInt(containerCpu, 10)
	w.Write([]byte(containerCpuStr))
}

type defaultHandler struct{}

func (h *defaultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
}

func main() {
	go cpu.Monitor()
	router := mux.NewRouter()
	router.Handle("/containers/{name}/mem", handlers.LoggingHandler(os.Stdout, &containerMemUsageHandler{}))
	router.Handle("/containers/{name}/cpu", handlers.LoggingHandler(os.Stdout, &containerCpuUsageHandler{}))
	router.Handle("/{misc}", handlers.LoggingHandler(os.Stdout, &defaultHandler{}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "4244"
	}
	fmt.Println("Listen on", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Println("Error bindint port:", err)
		os.Exit(1)
	}
}
