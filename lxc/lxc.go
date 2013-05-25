package lxc

import (
  "os/exec"
  "strings"
)

const (
  LXC_LIST_COMMAND = "lxc-ls"
  LXC_LIST_FLAG = "--active"
)

func ListContainer() ([]string, error) {
  output, err := exec.Command(LXC_LIST_COMMAND, LXC_LIST_FLAG).CombinedOutput()
  if(err != nil) {
    return nil, err
  }
  return strings.Split(string(output), "\n"), nil
}

