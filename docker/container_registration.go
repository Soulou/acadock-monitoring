package docker

import "sync"

var (
	registeredChans   []chan string
	registrationMutex *sync.Mutex
)

func init() {
	registrationMutex = &sync.Mutex{}
	containers := make(chan string)
	go ListenNewContainers(containers)
	go func() {
		for c := range containers {
			registrationMutex.Lock()
			for _, registeredChan := range registeredChans {
				registeredChan <- c
			}
			registrationMutex.Unlock()
		}
	}()
}

func RegisterToContainersStream() chan string {
	c := make(chan string, 1)
	registrationMutex.Lock()
	defer registrationMutex.Unlock()
	registeredChans = append(registeredChans, c)
	go ListRunningContainers(c)
	return c
}
