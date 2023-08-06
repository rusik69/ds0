package main

import (
	"fmt"

	"github.com/rusik69/ds0/pkg/client/cluster"
	"github.com/rusik69/ds0/pkg/client/cmdargs"
	"github.com/rusik69/ds0/pkg/client/file"
	"github.com/rusik69/ds0/pkg/client/node"
)

var version string

func main() {
	cmdargs.Instance = cmdargs.Parse()
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
		node, err := node.Stats(cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
		fmt.Println("Total space: ", node.TotalSpace)
		fmt.Println("Free space: ", node.FreeSpace)
		fmt.Println("Used space: ", node.UsedSpace)
	case "clusterstats":
		stats, err := cluster.Stats(cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
		fmt.Println("Total space: ", stats.TotalSpace)
		fmt.Println("Total free space: ", stats.TotalFreeSpace)
		fmt.Println("Total used space: ", stats.TotalUsedSpace)
		fmt.Println("Total nodes: ", stats.NodesCount)
		fmt.Println("Replicas: ", stats.Replicas)
	case "version":
		fmt.Println("Version: ", version)
	default:
		panic("unknown action: " + cmdargs.Instance.Cmd)
	}
}
