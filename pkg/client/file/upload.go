package file

import (
	"os"

	"github.com/rusik69/ds0/pkg/client/cmdargs"
)

// Upload uploads a file to the server.
func Upload() {
	host := cmdargs.CmdArgsInstance.HostName
	port := string(cmdargs.CmdArgsInstance.Port)
	src := cmdargs.CmdArgsInstance.Arg1
	dst := cmdargs.CmdArgsInstance.Arg2
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	url := "http://" + host + ":" + port + "/upload/" + dst
	resp, _ := http.Post(url, "application/octet-stream", file)
	defer resp.Body.Close()
}
