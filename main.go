package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	"github.com/zhenggu/wbx3ocp/builder/openstack"
	"github.com/zhenggu/wbx3ocp/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(openstack.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
