package cmdargs

import (
	"os"
	"strconv"
)

// Parse parses the command line arguments.
func Parse() CmdArgs {
	if len(os.Args) < 2 {
		panic("specify action: upload, download, addnode, removenode, listnodes, nodestats")
	}
	cmd := os.Args[1]
	var arg1, arg2, arg3 string
	switch cmd {
	case "upload":
		if len(os.Args) < 4 {
			panic("specify source file and destination path")
		}
		arg1 = os.Args[2]
		arg2 = os.Args[3]
	case "download":
		if len(os.Args) < 4 {
			panic("specify source path and destination file")

		}
		arg1 = os.Args[2]
		arg2 = os.Args[3]
	case "addnode":
		if len(os.Args) < 5 {
			panic("specify node name, hostname and port")
		}
		arg1 = os.Args[2]
		arg2 = os.Args[3]
		arg3 = os.Args[4]
	case "removenode":
		if len(os.Args) < 4 {
			panic("specify node name and port")
		}
		arg1 = os.Args[2]
	}
	portStr := strconv.Itoa(port)
	return CmdArgs{
		Cmd:      cmd,
		HostName: hostName,
		Port:     portStr,
		Arg1:     arg1,
		Arg2:     arg2,
		Arg3:     arg3,
	}
}
