package cmdargs

// CmdArgs is the command line arguments.
type CmdArgs struct {
	Cmd      string
	HostName string
	Port     string
	Arg1     string
	Arg2     string
	Arg3     string
}

// Instance is the singleton instance of CmdArgs.
var Instance CmdArgs
