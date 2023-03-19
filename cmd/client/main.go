package main

import (
	"github.com/rusik69/ds0/pkg/client/cmdargs"
	"github.com/rusik69/ds0/pkg/client/file"
	"github.com/rusik69/ds0/pkg/client/node"
)

func main() {
	cmdargs.Parse()
	if cmdargs.CmdArgsInstance.Cmd == "upload" {
		err := file.Upload(cmdargs.CmdArgsInstance.Arg1, cmdargs.CmdArgsInstance.Arg2, cmdargs.CmdArgsInstance.HostName, cmdargs.CmdArgsInstance.Port)
		if err != nil {
			panic(err)
		}
	} else if cmdargs.CmdArgsInstance.Cmd == "download" {
		err := file.Download(cmdargs.CmdArgsInstance.Arg1, cmdargs.CmdArgsInstance.Arg2, cmdargs.CmdArgsInstance.HostName, cmdargs.CmdArgsInstance.Port)
		if err != nil {
			panic(err)
		}
	} else if cmdargs.CmdArgsInstance.Cmd == "addnode" {
		err := node.Add(cmdargs.CmdArgsInstance.Arg1, cmdargs.CmdArgsInstance.Arg2, cmdargs.CmdArgsInstance.Arg3, cmdargs.CmdArgsInstance.HostName, cmdargs.CmdArgsInstance.Port)
		if err != nil {
			panic(err)
		}
	} else if cmdargs.CmdArgsInstance.Cmd == "removenode" {
		err := node.Remove(cmdargs.CmdArgsInstance.Arg1, cmdargs.CmdArgsInstance.HostName, cmdargs.CmdArgsInstance.Port)
		if err != nil {
			panic(err)
		}
	}
}
