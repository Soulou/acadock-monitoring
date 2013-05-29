package lxc

import (
  "os/exec"
  "strings"
)

const (
  LXC_LIST_COMMAND = "lxc-ls"
  LXC_LIST_FLAG = "--active"
)

func ListContainers() ([]string, error) {
  output, err := exec.Command(LXC_LIST_COMMAND, LXC_LIST_FLAG).CombinedOutput()
  if(err != nil) {
    return nil, err
  }
  if len(output) == 0 {
    return []string{}, nil
  }

  return strings.Split(string(output), "\n"), nil
}

