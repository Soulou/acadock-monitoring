package docker

import (
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/errgo.v1"
)

func ListenNewContainers(ids chan string) error {
	client, err := Client()
	if err != nil {
		return errgo.Mask(err)
	}

	listener := make(chan *docker.APIEvents)
	err = client.AddEventListener(listener)
	if err != nil {
		return errgo.Mask(err)
	}

	for event := range listener {
		if event.Status == "start" {
			ids <- event.ID
		}
	}
	return nil
}
