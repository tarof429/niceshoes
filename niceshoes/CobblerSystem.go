package niceshoes

import (
	"log"
	"os/exec"
)

const (
	CMD = "cobbler"
)

type CobblerSystem struct {
	Name          string       `json:"name"`
	Hostname      string       `json:"hostname"`
	Profile       string       `json:"profile"`
	KernelOptions string       `json:"kernelOptons"`
	NameServers   string       `json:"nameServers"`
	NextServerV4    string       `json:"nextServerV4"`
	NICs   []CobblerSystemNIC `json:"nics"`
}

func (c CobblerSystem) GetName() string {
	return c.Name
}

func (c CobblerSystem) GetHostname() string {
	return c.Hostname
}

func (c CobblerSystem) GetProfile() string {
	return c.Profile
}

func (c CobblerSystem) GetKernelOptions() string {
	return c.KernelOptions
}

func (c CobblerSystem) GetNameServers() string {
	return c.NameServers
}

func (c CobblerSystem) GetNextServerV4() string {
	return c.NextServerV4
}

// GetCmdline returns the command-line string bsed on the command and
// values in Cinterface
func (c CobblerSystem) GetCmdLine(command string, inter CobblerSystemNIC) []string {
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
				"--name-servers="+ "\"" + c.NameServers + "\"")
		}

		if len(c.NextServerV4) > 0 {
			cmdLine = append(cmdLine,
				"--next-server-v4="+c.NextServerV4)
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

func (c CobblerSystem) SystemExists() bool {

	args := c.GetCmdLine("list", CobblerSystemNIC{})

	cmdResult := exec.Command(CMD, args...)

	return cmdResult.Err != nil
}

func (c CobblerSystem) Import() error {

	if c.SystemExists() {
		return nil
	}

	//log.Println("Adding server")

	args := c.GetCmdLine("add", CobblerSystemNIC{})

	log.Printf("Running cobbler %s", args)

	cmd := exec.Command(CMD, args...)

	_, err := cmd.Output()

	if err != nil {
		//log.Printf("Error while adding: %s\n", c.Name)
		return err
	}

	// log.Println("Adding interfaces")

	for _, inter := range c.NICs {
		args := c.GetCmdLine("edit", inter)

		log.Printf("Running cobbler %s", args)

		cmd = exec.Command(CMD, args...)

		_, err = cmd.Output()

		if err != nil {
			log.Printf("%s\n", err.Error())
		}
	}

	args = c.GetCmdLine("remove-default-interface", CobblerSystemNIC{})

	cmd = exec.Command(CMD, args...)

	_, err = cmd.Output()

	if err != nil {
		log.Printf("Unable to remove default interface: %s\n", err.Error())
		return err
	}
	
	//log.Printf("Successfully added %s\n", c.Name)

	return nil
}