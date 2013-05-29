package mem

import (
  "github.com/Soulou/acadock-live-lxc/lxc"
  "testing"
)

func TestGetUsage(t *testing.T) {

   containers, err := lxc.ListContainers()
   if err != nil {
     t.Fatal(err)
   }
   if len(containers) == 0 {
     t.Log("There isn't any existing container, please create one")
     t.Log("Example : ")
     t.Log("\tlxc-start -d --name test_cpu_container /usr/bin/sleep 10")
     t.FailNow()
   }

  memUsage, err := GetUsage(containers[0])
  if err != nil {
    t.Fatal(err)
  }
  if memUsage < 0 {
    t.Fatal("Mem usage < 0")
  }

  t.Log("Mem usage of ", containers[0], " : ", memUsage)
}
