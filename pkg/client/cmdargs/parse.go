package cmdargs

import (
	"flag"
	"os"
)

// Parse parses the command line arguments.
func Parse() {
	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}
	cmd := os.Args[1]
	var arg1, arg2 string
	if cmd == "upload" {
		if len(os.Args) < 4 {
			flag.Usage()
			os.Exit(1)
		}
		arg1 = os.Args[2]
		arg2 = os.Args[3]
	}
	hostName := flag.String("host", "localhost", "host name")
	port := flag.Int("port", 6969, "port number")
	flag.Parse()
	CmdArgsInstance = &CmdArgs{
		HostName: *hostName,
		Port:     *port,
		Arg1:     arg1,
		Arg2:     arg2,
	}
}
