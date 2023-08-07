package niceshoes

import "log"

type SimulatorCobblerSystem struct {
	Name          string       `json:"name"`
	Hostname      string       `json:"hostname"`
	Profile       string       `json:"profile"`
	KernelOptions string       `json:"kernelOptons"`
	NameServers   string       `json:"nameServers"`
	NextServerV4    string       `json:"nextServerV4"`
	Cinterfaces   []Cinterface `json:"interfaces"`
}

func (c SimulatorCobblerSystem) GetName() string {
	return c.Name
}

func (c SimulatorCobblerSystem) GetHostname() string {
	return c.Hostname
}

func (c SimulatorCobblerSystem) GetProfile() string {
	return c.Profile
}

func (c SimulatorCobblerSystem) GetKernelOptions() string {
	return c.KernelOptions
}

func (c SimulatorCobblerSystem) GetNameServers() string {
	return c.NameServers
}

func (c SimulatorCobblerSystem) GetNextServerV4() string {
	return c.NextServerV4
}

func (c SimulatorCobblerSystem) GetCmdLine(command string, inter Cinterface) []string {

	realSystem := SystemCobblerSystem(c)

	return realSystem.GetCmdLine(command, inter)
}

func (c SimulatorCobblerSystem) Import() error {

	args := c.GetCmdLine("add", Cinterface{})

	log.Printf("Running %s %s", CMD, args)

	for _, inter := range c.Cinterfaces {
		args := c.GetCmdLine("edit", inter)

		log.Printf("Running %s %s", CMD, args)
	}

	return nil
}
