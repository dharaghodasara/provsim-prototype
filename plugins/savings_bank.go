package plugins

import (
	"fmt"

	"provsim-prototype/core"
)

//SavingsBank ...
type SavingsBank struct {
	core.PluginInfo
}

//PlugIt ...
func (s SavingsBank) PlugIt() {
	fmt.Println(s)
}
