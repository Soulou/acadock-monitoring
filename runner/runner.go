package runner

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/Scalingo/go-netns"
	"github.com/Scalingo/go-netstat"
)

func NetStatsRunner(pid string) (netstat.NetworkStats, error) {
	ns, err := netns.Setns(pid)
	if err != nil {
		return nil, err
	}
	defer ns.Close()

	path, err := exec.LookPath("net_stats_runner")
	if err != nil {
		return nil, err
	}
	stdout := new(bytes.Buffer)
	cmd := exec.Command(path)
	cmd.Stdout = stdout
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	var stats netstat.NetworkStats
	err = json.NewDecoder(stdout).Decode(&stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
