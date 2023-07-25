package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	cmd = "cobbler"
)

type Csystem struct {
	Name        string       `json:"name"`
	Hostname    string       `json:"hostname"`
	Profile     string       `json:"profile"`
	Cinterfaces []Cinterface `json:"interfaces"`
}

type Cinterface struct {
	Name            string `json:"name"`
	Netmask         string `json:"netmask"`
	MACAdress       string `json:"macaddress"`
	IPAddress       string `json:"ipAddress"`
	Gateway         string `json:"gateway"`
	DhcpTag         string `json:"dhcpTag"`
	InterfaceType   string `json:"interfaceType"`
	InterfaceMaster string `json:"interfaceMaster"`
	BondingOpts     string `json:"bondingOpts"`
	Static          string `json:"static"`
	KernelOptions   string `json:"kernelOptons"`
}

func (c Csystem) SystemExists() bool {

	args := []string{"system", "list", "--name=" + c.Name}

	cmdResult := exec.Command(cmd, args...)

	return cmdResult.Err != nil
}

func (c Csystem) Add(file *string) error {

	data, err := ioutil.ReadFile(*file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}

	if err := json.Unmarshal(data, &c); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s\n", err)
	}

	if c.SystemExists() {
		return nil
	}

	args := []string{"system", "add", "--name=" + c.Name,
		"--hostname=" + c.Hostname,
		"--profile=" + c.Profile}

	out, _ := exec.Command(cmd, args...).Output()

	if out != nil {
		log.Printf("%s\n", out)
	}

	for _, inter := range c.Cinterfaces {
		args := []string{
			"system", "edit", "--name=" + c.Name,
			"--interface=" + inter.Name,
			"--ip-address=" + inter.IPAddress,
			"--mac=" + inter.MACAdress,
			"--gateway=" + inter.Gateway,
			"--static=" + inter.Static,
			"--netmask=" + inter.Netmask,
		}

		out, _ = exec.Command(cmd, args...).Output()

		if out != nil {
			log.Printf("%s\n", out)
		}

		out, _ = exec.Command(cmd, args...).Output()

		if out != nil {
			log.Printf("%s\n", out)
		}
	}

	return nil
}

func runCobblerCmd() error {

	cmd := exec.Command("cobbler", "system", "help")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		//return err
	}

	fmt.Print(string(stdout))

	return nil
}

func main() {

	fmt.Println("Starting main...")

	log.SetOutput(os.Stdout)

	file := flag.String("file", "", "JSON file containg systems to import")

	flag.Parse()

	var c Csystem

	c.Add(file)

	// err := runCobblerCmd()

	// if err != nil {
	// 	log.Fatal(err)
	// }

}
