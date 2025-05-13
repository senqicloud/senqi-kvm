package dto

// VmCreateRequest 虚拟机创建请求参数结构体
type VmCreateRequest struct {
	Name      string `json:"name" binding:"required"`
	CPU       int    `json:"cpu" binding:"required"`
	Memory    int    `json:"memory" binding:"required"`
	Disk      int    `json:"disk" binding:"required"`
	ImageID   string `json:"image_id" binding:"required"`
	NetworkID string `json:"network_id" binding:"required"`
	Hostname  string `json:"hostname"`
	SSHKey    string `json:"ssh_key"`
	UserData  string `json:"user_data"`
	AutoStart bool   `json:"auto_start"`
}
