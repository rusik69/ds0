package cmdargs

import (
	"flag"
	"os"
	"strconv"
)

// Parse parses the command line arguments.
func Parse() {
	hostName := flag.String("host", "localhost", "host name")
	port := flag.Int("port", 6969, "port number")
	if len(os.Args) < 2 {
		panic("specify action: upload or download")
	}
	cmd := os.Args[1]
	var arg1, arg2, arg3 string
	if cmd == "upload" {
		if len(os.Args) < 4 {
			panic("specify source file and destination path")
		}
		arg1 = os.Args[2]
		arg2 = os.Args[3]
	} else if cmd == "download" {
		if len(os.Args) < 4 {
			panic("specify source path and destination file")

		}
		arg1 = os.Args[2]
		arg2 = os.Args[3]
	} else if cmd == "addnode" {
		if len(os.Args) < 5 {
			panic("specify node name, hostname and port")
		}
		arg1 = os.Args[2]
		arg2 = os.Args[3]
		arg3 = os.Args[4]
	} else if cmd == "removenode" {
		if len(os.Args) < 4 {
			panic("specify node name")
		}
		arg1 = os.Args[2]
	} else {
		panic("unknown action: " + cmd)
	}
	flag.Parse()
	portStr := strconv.Itoa(*port)
	CmdArgsInstance = &CmdArgs{
		Cmd:      cmd,
		HostName: *hostName,
		Port:     portStr,
		Arg1:     arg1,
		Arg2:     arg2,
		Arg3:     arg3,
	}
}
