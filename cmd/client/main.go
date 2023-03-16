package main

import (
	"github.com/rusik69/ds0/pkg/client/cmdargs"
	"github.com/rusik69/ds0/pkg/client/file"
)

func main() {
	cmdargs.Parse()
	if cmdargs.CmdArgsInstance.Cmd == "upload" {
		file.Upload(cmdargs.CmdArgsInstance.Arg1, cmdargs.CmdArgsInstance.Arg2, cmdargs.CmdArgsInstance.HostName, cmdargs.CmdArgsInstance.Port)
	} else if cmdargs.CmdArgsInstance.Cmd == "download" {
		file.Download(cmdargs.CmdArgsInstance.Arg1, cmdargs.CmdArgsInstance.Arg2, cmdargs.CmdArgsInstance.HostName, cmdargs.CmdArgsInstance.Port)
	}
}
