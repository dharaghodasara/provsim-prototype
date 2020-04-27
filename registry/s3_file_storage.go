package registry

import (
	"github.com/elangovans/provsim-prototype/core"
	"github.com/google/uuid"
)

//S3FileStorageProvider ...
type S3FileStorageProvider struct {
}

//Fetch ...
//TODO: Yet to be implemented
func (s S3FileStorageProvider) Fetch() (map[uuid.UUID]core.Provider, error) {
	var configs = make(map[uuid.UUID]core.Provider)

	return configs, nil
}
