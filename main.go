package main

import (
	log "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/nomad/plugins"
	"github.com/nicecode/nomad-driver-lxd/lxd"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of the LXC driver plugin
func factory(log log.Logger) interface{} {
	return lxd.NewLXDDriver(log)
}
