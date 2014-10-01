package mem

import (
	"log"
	"os"
	"strconv"

	"github.com/Soulou/acadock-monitoring/config"
	"github.com/Soulou/acadock-monitoring/docker"
)

const (
	LXC_MEM_USAGE_FILE = "memory.usage_in_bytes"
)

func GetUsage(id string) (int64, error) {
	id, err := docker.ExpandId(id)
	if err != nil {
		log.Println("Error when expanding id:", err)
		return 0, err
	}

	path := config.CgroupPath("memory", id) + "/" + LXC_MEM_USAGE_FILE
	f, err := os.Open(path)
	if err != nil {
		log.Println("Error while opening:", err)
		return 0, err
	}
	defer f.Close()

	buffer := make([]byte, 16)
	n, err := f.Read(buffer)
	if err != nil {
		log.Println("Error while reading ", path, ":", err)
		return 0, err
	}

	buffer = buffer[:n-1]
	val, err := strconv.ParseInt(string(buffer), 10, strconv.IntSize)
	if err != nil {
		log.Println("Error while parsing ", string(buffer), " : ", err)
		return 0, err
	}

	return val, nil
}
