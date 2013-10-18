package mem

import (
	"github.com/Soulou/acadock-live-lxc/lxc"
	"testing"
)

const (
	LXC_START_COMMAND = "lxc-start"
)

func TestGetUsage(t *testing.T) {
	// Need Root permission
	/* err := exec.Command(LXC_START_COMMAND, "-d", "--name", "test_cpu_container", "/usr/bin/sleep 10").Run() */
	/* if err != nil { */
	/*   t.Fatal("Please install", LXC_START_COMMAND, err) */
	/* } */

	containers, err := lxc.ListContainers()
	if err != nil {
		t.Fatal(err)
	}
	if len(containers) == 0 {
		t.Log("There isn't any existing container, please create one")
		t.Log("Example : ")
		t.Log("\t# lxc-start -d --name test_cpu_container /usr/bin/sleep 10")
		t.FailNow()
	}

	memUsage, err := GetUsage(containers[0])
	if err != nil {
		t.Fatal(err)
	}

	if memUsage < 0 || memUsage > 100 {
		t.Fatal("CPU usage out of [0,100]")
	}
}
