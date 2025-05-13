package internal

import "github.com/libvirt/libvirt-go"

// Conn libvirt 连接信息
var Conn *libvirt.Connect

// ConnError 连接异常信息
var ConnError error

func Setup() {
	Conn, ConnError = libvirt.NewConnect("qemu:///system")
}
