package config

import (
	"os"
)

var ENV = map[string]string{
	"DOCKER_URL":    "http://localhost:4243",
	"PORT":          "4244",
	"REFRESH_TIME":  "5",
	"CGROUP_SOURCE": "docker",
}

func init() {
	for k, v := range ENV {
		if os.Getenv(k) != "" {
			ENV[k] = os.Getenv(k)
		} else {
			os.Setenv(k, v)
		}
	}
}

func CgroupPath(cgroup string, id string) string {
	if ENV["CGROUP_SOURCE"] == "docker" {
		return "/sys/fs/cgroup/" + cgroup + "/docker/" + id
	} else if ENV["CGROUP_SOURCE"] == "systemd" {
		return "/sys/fs/cgroup/" + cgroup + "/system.slice/docker-" + id + ".scope"
	} else {
		panic("unknown cgroup source" + ENV["CGROUP_SOURCE"])
	}
}
