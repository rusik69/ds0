package server

// NodeStats is the info of a node.
type NodeStats struct {
	Host       string
	Port       string
	TotalSpace uint64
	FreeSpace  uint64
	UsedSpace  uint64
}
