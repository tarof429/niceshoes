package niceshoes

import (
	"log"
	"os/exec"
)

const (
	CMD = "cobbler"
)

type SystemCobblerSystem struct {
	Name          string       `json:"name"`
	Hostname      string       `json:"hostname"`
	Profile       string       `json:"profile"`
	KernelOptions string       `json:"kernelOptons"`
	NameServers   string       `json:"nameServers"`
	NextServer    string       `json:"nextServer"`
	Cinterfaces   []Cinterface `json:"interfaces"`
}

func (c SystemCobblerSystem) GetName() string {
	return c.Name
}

func (c SystemCobblerSystem) GetHostname() string {
	return c.Hostname
}

func (c SystemCobblerSystem) GetProfile() string {
	return c.Profile
}

func (c SystemCobblerSystem) GetKernelOptions() string {
	return c.KernelOptions
}

func (c SystemCobblerSystem) GetNameServers() string {
	return c.NameServers
}

func (c SystemCobblerSystem) GetNextServer() string {
	return c.NextServer
}

// GetCmdline returns the command-line string bsed on the command and
// values in Cinterface
func (c SystemCobblerSystem) GetCmdLine(command string, inter Cinterface) []string {
	var cmdLine []string

	if command == "list" {
		cmdLine = []string{"system", "list", "--name=" + c.Name}
	} else if command == "add" {
		cmdLine = []string{"system", "add", "--name=" + c.Name,
			"--hostname=" + c.Hostname,
			"--profile=" + c.Profile}

		if len(c.KernelOptions) > 0 {
			cmdLine = append(cmdLine,
				"--kernel-options="+c.KernelOptions)
		}

		if len(c.NameServers) > 0 {
			cmdLine = append(cmdLine,
				"--name-servers="+c.NameServers)
		}

		if len(c.NextServer) > 0 {
			cmdLine = append(cmdLine,
				"--next-server="+c.NextServer)
		}

	} else if command == "edit" {
		cmdLine = []string{"system", "edit", "--name=" + c.Name,
			"--interface=" + inter.Name,
			"--ip-address=" + inter.IPAddress,
			"--mac=" + inter.MACAdress}

		if inter.Static == "1" {
			cmdLine = append(cmdLine,
				"--static="+inter.Static,
				"--gateway="+inter.Gateway,
				"--netmask="+inter.Netmask)
		}

		if len(inter.InterfaceMaster) > 0 {
			cmdLine = append(cmdLine,
				"--interface-master="+inter.InterfaceMaster)
		}

		if len(inter.InterfaceType) > 0 {
			cmdLine = append(cmdLine,
				"--interface-type="+inter.InterfaceType)
		}

		if len(inter.BondingOpts) > 0 {
			cmdLine = append(cmdLine,
				"--bonding-opts="+inter.BondingOpts)
		}

	} else if command == "remove-default-interface" {
		cmdLine = []string{"system", "edit", "--name=" + c.Name,
			"--delete-interface",
			"--interface=default"}
	}
	return cmdLine
}

func (c SystemCobblerSystem) SystemExists() bool {

	args := c.GetCmdLine("list", Cinterface{})

	cmdResult := exec.Command(CMD, args...)

	return cmdResult.Err != nil
}

func (c SystemCobblerSystem) Import() error {

	if c.SystemExists() {
		return nil
	}

	log.Println("Adding server")

	args := c.GetCmdLine("add", Cinterface{})

	// log.Printf("Running cobbler %s", args)

	cmd := exec.Command(CMD, args...)

	_, err := cmd.Output()

	if err != nil {
		log.Printf("Error while adding: %s\n", err.Error())
	}

	// log.Println("Adding interfaces")

	for _, inter := range c.Cinterfaces {
		args := c.GetCmdLine("edit", inter)

		cmd = exec.Command(CMD, args...)

		_, err = cmd.Output()

		if err != nil {
			log.Printf("%s\n", err.Error())
		}
	}

	args = c.GetCmdLine("remove-default-interface", Cinterface{})

	cmd = exec.Command(CMD, args...)

	_, err = cmd.Output()

	if err != nil {
		log.Printf("Unable to remove default interface: %s\n", err.Error())
		return err
	}
	
	log.Printf("Successfully added %s\n", c.Name)

	return nil
}