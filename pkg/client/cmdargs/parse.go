package cmdargs

import (
	"flag"
	"strconv"
)

// Parse parses the command line arguments.
func Parse() {
	hostName := flag.String("host", "localhost", "host name")
	port := flag.Int("port", 6969, "port number")
	if flag.NArg() < 2 {
		panic("specify action: upload or download")
	}
	cmd := flag.Arg(1)
	var arg1, arg2 string
	if cmd == "upload" {
		if flag.NArg() < 4 {
			panic("specify source file and destination path")
		}
		arg1 = flag.Arg(2)
		arg2 = flag.Arg(3)
	} else if cmd == "download" {
		if flag.NArg() < 4 {
			panic("specify source path and destination file")

		}
		arg1 = flag.Arg(2)
		arg2 = flag.Arg(3)
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
	}
}
