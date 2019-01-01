package test

import "fmt"

type ConfigWithPointers struct {
	Environment *string // pointer to string
	Version     *string
	HostName    *string
}

func (c *ConfigWithPointers) String() string {
	var envOut, verOut, hostOut string
	envOut = "<nil>"
	verOut = "<nil>"
	hostOut = "<nil>"

	if c.Environment != nil { // Check for nil!
		envOut = *c.Environment
	}

	if c.Version != nil {
		verOut = *c.Version
	}

	if c.HostName != nil {
		hostOut = *c.HostName
	}
	return fmt.Sprintf("Environment: '%v'\nVersion:'%v'\nHostName: '%v'",
		envOut, verOut, hostOut)
}