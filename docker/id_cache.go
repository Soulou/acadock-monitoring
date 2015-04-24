package docker

import (
	"errors"
	"strings"
)

var IdNotInCache = errors.New("id not found in cache")

var containerIdsCache = []string{}

func expandIdFromCache(id string) (string, error) {
	for _, fullId := range containerIdsCache {
		if strings.HasPrefix(fullId, id) {
			return fullId, nil
		}
	}
	return "", IdNotInCache
}
