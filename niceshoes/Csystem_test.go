package niceshoes

import (
	"fmt"
	"testing"
)

func compareArrays(x []string, y []string) bool {

	fmt.Println(x)

	fmt.Println(y)

	if len(x) != len(y) {
		return false
	}

	if len(x) <= 0 {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestList(t *testing.T) {

	c := SystemCobblerSystem{
		Name:        "test",
		Profile:     "test",
		Hostname:    "localhost",
		Cinterfaces: []Cinterface{},
	}

	cmdLine := c.GetCmdLine("list", Cinterface{})
	expected := []string{"system", "list", "--name=" + c.Name}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}
	t.Log("Pass")
}

func TestAdd(t *testing.T) {

	c := SystemCobblerSystem{
		Name:     "test",
		Profile:  "test",
		Hostname: "localhost",
	}

	cmdLine := c.GetCmdLine("add", Cinterface{})
	expected := []string{"system", "add", "--name=" +
		c.Name, "--hostname=" + c.Hostname, "--profile=" + c.Profile}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}
	t.Log("Pass")
}

func TestEdit(t *testing.T) {

	c := SystemCobblerSystem{
		Name:     "test",
		Profile:  "test",
		Hostname: "localhost",
		KernelOptions: "",
		Cinterfaces: []Cinterface{
			{
				Name:          "eth0",
				Netmask:       "255.255.255.0",
				MACAdress:     "AC:AB:AA:12:24:CA",
				Static:        "1",
				Gateway:       "192.168.1.1",
				IPAddress:     "192.168.1.10",

			},
		},
	}

	inter := c.Cinterfaces[0]

	cmdLine := c.GetCmdLine("edit", inter)
	expected := []string{"system", "edit", "--name=" + c.Name,
		"--interface=" + inter.Name,
		"--ip-address=" + inter.IPAddress,
		"--mac=" + inter.MACAdress,
		"--static=" + inter.Static,
		"--gateway=" + inter.Gateway,
		"--netmask=" + inter.Netmask}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}

	t.Log("Pass")
}

func TestRemoveDefaultInterface(t *testing.T) {

	c := SystemCobblerSystem{
		Name:        "test",
		Profile:     "test",
		Hostname:    "localhost",
		Cinterfaces: []Cinterface{},
	}

	cmdLine := c.GetCmdLine("remove-default-interface", Cinterface{})

	expected := []string{"system", "edit", "--name=" + c.Name,
		"--delete-interface",
		"--interface=default"}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}

	t.Log("Pass")
}

func TestSimulator(t *testing.T) {

	var c CobblerDomain

	c = SimulatorCobblerSystem{
		Name:     "test",
		Profile:  "test",
		Hostname: "localhost",
		KernelOptions: "\"foo=x bar=y\"",
		NameServers: "192.168.1.1",
		NextServer: "localhost",
		Cinterfaces: []Cinterface{},
	}

	cmdLine := c.GetCmdLine("add", Cinterface{})
	expected := []string{"system", "add", 
		"--name=" + c.GetName(), 
		"--hostname=" + c.GetHostname(), 
		"--profile=" + c.GetProfile(),
		"--kernel-options=" + c.GetKernelOptions(),
		"--name-servers=" + c.GetNameServers(),
		"--next-server=" + c.GetNextServer(),
	}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}
	t.Log("Pass")
}