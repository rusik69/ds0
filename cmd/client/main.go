package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
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
	case "listfiles":
		files, err := file.List(cmdargs.Instance.HostName, cmdargs.Instance.Port, cmdargs.Instance.Arg1)
		if err != nil {
			panic(err)
		}
		for fileName, fileInfo := range files {
			fmt.Println(fileName + " " + humanize.Bytes(fileInfo.Size))
		}
	case "nodestats":
		node, err := node.Stats(cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
		fmt.Println("Total space: ", humanize.Bytes(node.TotalSpace))
		fmt.Println("Free space: ", humanize.Bytes(node.FreeSpace))
		fmt.Println("Used space: ", humanize.Bytes(node.UsedSpace))
	case "clusterstats":
		stats, err := cluster.Stats(cmdargs.Instance.HostName, cmdargs.Instance.Port)
		if err != nil {
			panic(err)
		}
		fmt.Println("Total space: ", humanize.Bytes(stats.TotalSpace))
		fmt.Println("Total free space: ", humanize.Bytes(stats.TotalFreeSpace))
		fmt.Println("Total used space: ", humanize.Bytes(stats.TotalUsedSpace))
		fmt.Println("Total nodes: ", stats.NodesCount)
		fmt.Println("Replicas: ", stats.Replicas)
	case "version":
		fmt.Println("Version: ", version)
	default:
		panic("unknown action: " + cmdargs.Instance.Cmd)
	}
}
