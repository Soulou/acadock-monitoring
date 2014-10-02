package cpu

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Soulou/acadock-monitoring/config"
	"github.com/Soulou/acadock-monitoring/docker"
)

const (
	LXC_CPUACCT_USAGE_FILE = "cpuacct.usage"
)

var RefreshTime int

func init() {
	var err error
	RefreshTime, err = strconv.Atoi(config.ENV["REFRESH_TIME"])
	if err != nil {
		panic(err)
	}
}

var (
	previousCpuUsages = make(map[string]int64)
	cpuUsages         = make(map[string]int64)
)

func cpuacctUsage(container string) (int64, error) {
	file := config.CgroupPath("cpuacct", container) + "/" + LXC_CPUACCT_USAGE_FILE
	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	buffer := make([]byte, 16)
	n, err := f.Read(buffer)
	buffer = buffer[:n]

	bufferStr := string(buffer)
	bufferStr = bufferStr[:len(bufferStr)-1]

	res, err := strconv.ParseInt(bufferStr, 10, 64)
	if err != nil {
		log.Println("Failed to parse : ", err)
		return 0, err
	}
	return res, nil
}

func Monitor() {
	containers := make(chan string)
	go docker.ListenNewContainers(containers)
	go docker.ListRunningContainers(containers)
	for c := range containers {
		go monitorContainer(c)
	}
}

func monitorContainer(id string) {
	fmt.Println("monitor cpu", id)
	tick := time.NewTicker(time.Duration(RefreshTime) * time.Second)
	for {
		<-tick.C
		usage, err := cpuacctUsage(id)
		if err != nil {
			if _, ok := cpuUsages[id]; ok {
				delete(cpuUsages, id)
			}
			if _, ok := previousCpuUsages[id]; ok {
				delete(previousCpuUsages, id)
			}
			log.Println("stop monitoring", id)
			return
		}

		if prevUsage, ok := cpuUsages[id]; ok {
			previousCpuUsages[id] = prevUsage
		}
		cpuUsages[id] = usage
	}
}

func GetUsage(id string) (int64, error) {
	id, err := docker.ExpandId(id)
	if err != nil {
		log.Println("Error when expanding id:", err)
		return -1, err
	}
	if _, ok := previousCpuUsages[id]; !ok {
		return -1, errors.New("not ready")
	}
	return int64((float64((cpuUsages[id] - previousCpuUsages[id])) / float64(1e9) / float64(RefreshTime)) * 100), nil
}
