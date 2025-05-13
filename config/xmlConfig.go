package config

import (
	"fmt"
	"senqi-kvm/entity/dto"
)

func GetDefineVmXML(config dto.VmCreateRequest) string {
	// 构建 XML 模板并替换相应的字段
	return fmt.Sprintf(`
		<domain type='kvm'>
		  <name>%s</name>
		  <memory unit='MiB'>%d</memory>
		  <vcpu placement='static'>%d</vcpu>
		  <os>
			<type arch='x86_64'>hvm</type>
		  </os>
		  <devices>
			<disk type='file' device='disk'>
			  <driver name='qemu' type='qcow2'/>
			  <source file='%s'/>
			  <target dev='vda' bus='virtio'/>
			</disk>
			<interface type='network'>
			  <source network='%s'/>
			  <model type='virtio'/>
			</interface>
			<graphics type='vnc' port='-1'/>
		  </devices>
		  <metadata>
			<cloud-init>%s</cloud-init>
		  </metadata>
		  <features>
			<acpi/>
			<apic/>
		  </features>
		  <on_poweroff>destroy</on_poweroff>
		  <on_reboot>restart</on_reboot>
		  <on_crash>restart</on_crash>
		</domain>`,
		config.Name,
		config.Memory,
		config.CPU,
		config.DiskPath,
		config.Network,
		config.UserData,
	)
}
