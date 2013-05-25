package mem

import (
  "fmt"
  "log"
  "os"
)

const (
	LXC_MEM_DIR            = "/sys/fs/cgroup/memory/lxc"
	LXC_MEM_USAGE_FILE     = "memory.usage_in_bytes"
)

func GetUsage(name string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s/%s", LXC_MEM_DIR, name, LXC_MEM_USAGE_FILE)
	f, err := os.Open(path)
  defer f.Close()
	if err != nil {
		log.Println("Error while opening : ", err)
		return nil, err
	}

  buffer := make([]byte, 16)
  _, err = f.Read(buffer)
  if err != nil {
    log.Println("Error while reading ", path, " : ", err)
    return nil, err
  }

  return buffer, nil
}
