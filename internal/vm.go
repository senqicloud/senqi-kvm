package internal

import (
	"log"
	"senqi-kvm/config"
	"senqi-kvm/entity/dto"
)

type VmService interface {
	CreateVm(req dto.VmCreateRequest) (string, error)
}

type VmServiceImpl struct{}

// CreateVm 创建虚拟机
func (vm *VmServiceImpl) CreateVm(req dto.VmCreateRequest) (string, error) {
	createXML := config.GetDefineVmXML(req)

	// 定义虚拟机（不启动）
	domain, err := Conn.DomainDefineXML(createXML)

	if err != nil {
		log.Fatalf("创建虚拟机失败: %v", err)
	}
	defer domain.Free()

	// 启动虚拟机
	if err := domain.Create(); err != nil {
		log.Fatalf("启动虚拟机失败: %v", err)
	}

	log.Printf("虚拟机 %s 创建并启动成功", "vmName")

	return "创建成功!", nil
}
