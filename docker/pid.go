package docker

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/Soulou/acadock-monitoring/config"
)

func Pid(id string) (string, error) {
	time.Sleep(time.Second)
	path := config.CgroupPath("memory", id)
	content, err := ioutil.ReadFile(path + "/tasks")
	if err != nil {
		return "", err
	}
	return strings.Split(string(content), "\n")[0], nil
}
