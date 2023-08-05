package main

import (
	"flag"
	"fmt"

	"github.com/rusik69/ds0/pkg/client/cmdargs"
	"github.com/rusik69/ds0/pkg/client/file"
	"github.com/rusik69/ds0/pkg/client/node"
)

var version string

func main() {
	var hostName string
	var port int
	flag.StringVar(&hostName, "host", "localhost", "ns hostname")
	flag.IntVar(&port, "port", 6969, "port number")
	flag.Parse()
	cmdargs.Instance = cmdargs.Parse()
	cmdargs.Instance.HostName = hostName
	cmdargs.Instance.Port = fmt.Sprintf("%d", port)
	fmt.Println("Cmd: ", cmdargs.Instance.Cmd)
	fmt.Println("Host: ", cmdargs.Instance.HostName)
	fmt.Println("Port: ", cmdargs.Instance.Port)
	fmt.Println("Arg1: ", cmdargs.Instance.Arg1)
	fmt.Println("Arg2: ", cmdargs.Instance.Arg2)
	fmt.Println("Arg3: ", cmdargs.Instance.Arg3)
	switch cmdargs.Instance.Cmd {
	case "upload":
		err := file.Upload(cmdargs.Instance.Arg1, cmdargs.Instance.Arg2, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	case "download":
		err := file.Download(cmdargs.Instance.Arg1, cmdargs.Instance.Arg2, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	case "addnode":
		err := node.Add(cmdargs.Instance.Arg1, cmdargs.Instance.Arg2, cmdargs.Instance.Arg3, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	case "removenode":
		err := node.Remove(cmdargs.Instance.Arg1, cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
	case "listnodes":
		nodes, err := node.List(cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			fmt.Println(node.Host + ":" + node.Port)
		}
	case "nodestats":
		node, err := node.Stats(cmdargs.Instance.Arg1, cmdargs.Instance.Arg2)
		if err != nil {
			panic(err)
		}
		fmt.Println("Total space: ", node.TotalSpace)
		fmt.Println("Free space: ", node.FreeSpace)
		fmt.Println("Used space: ", node.UsedSpace)
	case "version":
		fmt.Println("Version: ", version)
	default:
		panic("unknown action: " + cmdargs.Instance.Cmd)
	}
}
