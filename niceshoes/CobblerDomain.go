package niceshoes

type CobblerDomain interface {
	Import() error
	GetName() string
	GetHostname() string
	GetProfile() string
	GetKernelOptions() string
	GetNameServers() string
	GetNextServerV4() string
	GetCmdLine(command string, inter Cinterface) []string
}

