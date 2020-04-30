package registry

import (
	"provsim-prototype/core"

	"github.com/google/uuid"
)

//RegistryMap ...
var RegistryMap map[uuid.UUID]core.Provider

//RegisryLoaded ...
var RegisryLoaded error
