package cpu

import (
	"github.com/Soulou/acadock-live-lxc/lxc"
	"fmt"
	"github.com/Soulou/acadock-live-lxc/lxc/utils"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	LXC_CPUACCT_DIR        = "/sys/fs/cgroup/cpuacct/lxc"
	LXC_CPUACCT_USAGE_FILE = "cpuacct.usage"
	REFRESH_TIME           = 1
)

var (
	previousCpuUsages = make(map[string]int64)
	cpuUsages         = make(map[string]int64)
)

func cpuacctUsage(container string) (int64, error) {
	file := fmt.Sprintf("%s/%s/%s", LXC_CPUACCT_DIR, container, LXC_CPUACCT_USAGE_FILE)
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

	res, err := strconv.ParseInt(bufferStr, 10, strconv.IntSize)
	if err != nil {
		log.Println("Failed to parse : ", err)
		return 0, err
	}
	return res, nil
}

func Monitor() {
	tick := time.NewTicker(REFRESH_TIME * time.Second)
	for {
		<-tick.C
		for k, v := range cpuUsages {
			previousCpuUsages[k] = v
		}

		containers, err := lxc.ListContainers()
		if err != nil {
			log.Fatalln(err)
		}

		for _, container := range containers {
			n, err := cpuacctUsage(container)
			if err != nil {
				log.Fatalln(err)
			}
			cpuUsages[container] = n
		}
	}
}

func GetUsage(id_short string) int64 {
	id , err := utils.GetFullContainerId(id_short)
	if err != nil {
		return -1
	}
	return int64((float64((cpuUsages[id] - previousCpuUsages[id])) / float64(1e9) / float64(runtime.NumCPU())) * 100)
}
