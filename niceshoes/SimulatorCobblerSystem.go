package niceshoes

import "log"

type SimulatorCobblerSystem struct {
	Name          string       `json:"name"`
	Hostname      string       `json:"hostname"`
	Profile       string       `json:"profile"`
	KernelOptions string       `json:"kernelOptons"`
	NameServers   string       `json:"nameServers"`
	NextServerV4    string       `json:"nextServerV4"`
	Cinterfaces   []CobblerSystemNIC `json:"interfaces"`
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

func (c SimulatorCobblerSystem) GetCmdLine(command string, inter CobblerSystemNIC) []string {

	temp := CobblerSystem{
		Hostname:  c.Hostname,
		KernelOptions: c.KernelOptions,
		Name:  c.Name,
		NameServers: c.NameServers,
		NextServerV4: c.NextServerV4,
		Profile:  c.Profile,
	}

	return temp.GetCmdLine(command, inter)
}

func (c SimulatorCobblerSystem) Import() error {

	args := c.GetCmdLine("add", CobblerSystemNIC{})

	log.Printf("Running %s %s", CMD, args)

	for _, inter := range c.Cinterfaces {
		args := c.GetCmdLine("edit", inter)

		log.Printf("Running %s %s", CMD, args)
	}

	return nil
}
