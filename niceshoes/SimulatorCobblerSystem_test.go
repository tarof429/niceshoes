package niceshoes

import "testing"

func TestSimulator(t *testing.T) {

	var c CobblerDomain

	c = SimulatorCobblerSystem{
		Name:     "test",
		Profile:  "test",
		Hostname: "localhost",
		KernelOptions: "\"foo=x bar=y\"",
		NameServers: "192.168.1.1",
		NextServerV4: "localhost",
		Cinterfaces: []Cinterface{},
	}

	cmdLine := c.GetCmdLine("add", Cinterface{})
	expected := []string{"system", "add", 
		"--name=" + c.GetName(), 
		"--hostname=" + c.GetHostname(), 
		"--profile=" + c.GetProfile(),
		"--kernel-options=" + c.GetKernelOptions(),
		"--name-servers=\"" + c.GetNameServers() + "\"",
		"--next-server=" + c.GetNextServerV4(),
	}

	if compareArrays(cmdLine, expected) != true {
		t.Fail()
	}
	t.Log("Pass")
}