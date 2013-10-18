package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	LXC_DIR = "/sys/fs/cgroup/memory/lxc"
)

func GetFullContainerId(name string) (string, error) {
	f, err := os.Open(LXC_DIR)
	if err != nil {
		return "", fmt.Errorf("Unable to open LXC cgroup dir: %v", err)
	}
	names, err := f.Readdirnames(0)
	if err != nil {
		return "", fmt.Errorf("Unable to list directories in %s: %v", LXC_DIR, err)
	}

	validNames := make([]string, 0)
	for _, n := range names {
		if strings.HasPrefix(n, name) {
			validNames = append(validNames, n)
		}
	}
	if len(validNames) != 1 {
		return "", fmt.Errorf("Invalid container id %s", name)
	}

	return validNames[0], nil
}
