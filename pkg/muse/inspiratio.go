package muse

import consul "github.com/log-ed/orazio/pkg/consul"

func GetInspiratio() []string {
	// Connect to consul api, make sure to have consul agent running on the default port
	return consul.GetMuse()
}
