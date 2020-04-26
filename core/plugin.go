package core

//PluginInfo ...
type PluginInfo struct {
	Name string
	Desc string
}

//Plugin ...
type Plugin interface {
	PlugIt()
}
