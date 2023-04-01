package main

import (
	"github.com/rusik69/ds0/pkg/client/cmdargs"
	"github.com/rusik69/ds0/pkg/client/file"
	"github.com/rusik69/ds0/pkg/client/node"
)

func main() {
	cmdargs.Parse()
	if cmdargs.Instance.Cmd == "upload" {
		err := file.Upload(cmdargs.Instance.Arg1, cmdargs.Instance.Arg2, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	} else if cmdargs.Instance.Cmd == "download" {
		err := file.Download(cmdargs.Instance.Arg1, cmdargs.Instance.Arg2, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	} else if cmdargs.Instance.Cmd == "addnode" {
		err := node.Add(cmdargs.Instance.Arg1, cmdargs.Instance.Arg2, cmdargs.Instance.Arg3, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	} else if cmdargs.Instance.Cmd == "removenode" {
		err := node.Remove(cmdargs.Instance.Arg1, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	}
}
