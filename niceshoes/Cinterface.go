package niceshoes

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
}


