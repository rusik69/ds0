package cmdargs

// CmdArgs is the command line arguments.
type CmdArgs struct {
	HostName string
	Port     int
	Arg1     string
	Arg2     string
}

// CmdArgsInstance is the singleton instance of CmdArgs.
var CmdArgsInstance *CmdArgs
