package docker

import (
	"github.com/Soulou/acadock-monitoring/config"
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/errgo.v1"
)

func Client() (*docker.Client, error) {
	client, err := docker.NewClient(config.ENV["DOCKER_URL"])
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return client, nil
}
