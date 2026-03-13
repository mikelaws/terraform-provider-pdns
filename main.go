package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mikelaws/terraform-provider-pdns/powerdns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: powerdns.Provider})
}
