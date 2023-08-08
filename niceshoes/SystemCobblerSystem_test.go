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

func TestReport(t *testing.T) {

	c := CobblerSystem{
		Name:        "test",
		Profile:     "test",
		Hostname:    "localhost",
		NICs: []CobblerSystemNIC{},
	}

	cmdLine := c.GetCmdLine("report", CobblerSystemNIC{})
	expected := []string{"system", "report", "--name=" + c.Name}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}
	t.Log("Pass")
}

func TestAdd(t *testing.T) {

	c := CobblerSystem{
		Name:     "test",
		Profile:  "test",
		Hostname: "localhost",
	}

	cmdLine := c.GetCmdLine("add", CobblerSystemNIC{})
	expected := []string{"system", "add", "--name=" +
		c.Name, "--hostname=" + c.Hostname, "--profile=" + c.Profile}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}
	t.Log("Pass")
}

func TestEdit(t *testing.T) {

	c := CobblerSystem{
		Name:     "test",
		Profile:  "test",
		Hostname: "localhost",
		KernelOptions: "",
		NICs: []CobblerSystemNIC{
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

	inter := c.NICs[0]

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

	c := CobblerSystem{
		Name:        "test",
		Profile:     "test",
		Hostname:    "localhost",
		NICs: []CobblerSystemNIC{},
	}

	cmdLine := c.GetCmdLine("remove-default-interface", CobblerSystemNIC{})

	expected := []string{"system", "edit", "--name=" + c.Name,
		"--delete-interface",
		"--interface=default"}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}

	t.Log("Pass")
}