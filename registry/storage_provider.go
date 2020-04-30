package registry

import (
	"provsim-prototype/core"

	"github.com/google/uuid"
)

//StorageProvider is an interface to get registry entiries from
// for example, local file system or S3 folder or database
// or any other storage mechanism
type StorageProvider interface {
	Fetch() (map[uuid.UUID]core.Provider, error)
}
