package cmdargs

import (
	"flag"
	"fmt"
)

// Parse parses the command line arguments.
func Parse() CmdArgs {
	var hostName string
	var port int
	flag.StringVar(&hostName, "host", "localhost", "ns hostname")
	flag.IntVar(&port, "port", 6969, "port number")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		panic("specify action: upload, download, addnode, removenode, listnodes, nodestats, clusterstats")
	}
	cmd := args[0]
	var arg1, arg2, arg3 string
	switch cmd {
	case "upload":
		if len(args) < 3 {
			panic("specify source file and destination path")
		}
		arg1 = args[1]
		arg2 = args[2]
	case "download":
		if len(args) < 3 {
			panic("specify source path and destination file")

		}
		arg1 = args[1]
		arg2 = args[2]
	case "addnode":
		if len(args) < 4 {
			panic("specify node name, hostname and port")
		}
		arg1 = args[1]
		arg2 = args[2]
		arg3 = args[3]
	case "removenode":
		if len(args) < 3 {
			panic("specify node name and port")
		}
		arg1 = args[1]
		arg2 = args[2]
	}
	return CmdArgs{
		Cmd:      cmd,
		Arg1:     arg1,
		Arg2:     arg2,
		Arg3:     arg3,
		HostName: hostName,
		Port:     fmt.Sprintf("%d", port),
	}
}
