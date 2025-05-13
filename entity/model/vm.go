package model

type Vm struct {
	Name      string `json:"name"`       // 虚拟机名称
	UUID      string `json:"uuid"`       // 虚拟机 UUID
	State     string `json:"state"`      // 当前状态（运行中、关闭等）
	CPU       uint   `json:"cpu"`        // 分配的虚拟 CPU 数
	Memory    uint64 `json:"memory"`     // 分配的内存 (单位：KiB)
	MaxMemory uint64 `json:"max_memory"` // 最大内存 (单位：KiB)
	DiskPath  string `json:"disk_path"`  // 镜像文件路径
	Network   string `json:"network"`    // 网络名称或桥接名
	IPAddress string `json:"ip_address"` // IP 地址（需启用 guest agent 才能获取）
}

// VmConfig 虚拟机配置结构体
type VmConfig struct {
	Name      string
	Memory    int    // MB
	CPU       int    // 核心数
	DiskPath  string // 磁盘文件路径
	Network   string // 网络配置
	SSHKey    string // SSH 公钥
	UserData  string // Cloud-Init 配置
	AutoStart bool   // 是否自动启动
}
