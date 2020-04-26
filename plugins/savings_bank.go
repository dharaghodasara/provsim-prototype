package plugins

import (
	"fmt"

	"github.com/elangovans/provsim-prototype/core"
)

//SavingsBank ...
type SavingsBank struct {
	core.PluginInfo
}

//PlugIt ...
func (s SavingsBank) PlugIt() {
	fmt.Println(s)
}
