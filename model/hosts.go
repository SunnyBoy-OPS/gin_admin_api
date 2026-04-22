package model

// ConnectTestHost 连接主机的接口
type ConnectTestHost struct {
	Host     string `json:"host" binding:"required"`
	Port     int    `json:"port" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// HostsInfo 实现连接主机信息接口
type HostsInfo struct {
	Hosts    string `json:"hosts"`
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	Group    string `json:"group"`
	OS       string `json:"os"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

type CmdInfo struct {
	CpuCore   int    `json:"CpuCore"`
	MemTotal  string `json:"MemTotal"`
	DiskTotal string `json:"DiskTotal"`
}

// CpuInfo 实现主机cpu信息接口
type CpuInfo struct {
	CpuCore  int    `json:"CpuCore"`
	CpuTotal uint64 `json:"CpuTotal"`
	CpuUsage uint64 `json:"CpuUsage"`
}

// MemInfo 实现主机内存信息接口
type MemInfo struct {
	MemTotal int    `json:"MemTotal"`
	MemUsage uint64 `json:"MemUsage"`
}

// DiskInfo 实现主机磁盘信息接口
type DiskInfo struct {
	DiskTotal int    `json:"DiskTotal"`
	DiskUsage uint64 `json:"DiskUsage"`
}
