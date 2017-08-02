package main

import (
	"fmt"
	"time"

	"../ovirtsdk4"
)

func main() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %s\n", err.Error())
	}
	defer conn.Close()

	// Get the reference to the "vms" service:
	vmsService := conn.SystemService().VmsService()

	// Use the "list" method of the "vms" service to list all the virtual machines of the system:
	vmsResponse, err := vmsService.List().Send()

	if err != nil {
		fmt.Printf("Failed to get vm list, reason: %v\n", err)
		return
	}

	// Print the virtual machine names and identifiers:
	for _, vm := range vmsResponse.Vms() {
		fmt.Printf("VM - (name: %v, id: %v)\n", *vm.Name, *vm.Id)
		fmt.Printf("VM devices: %+v\n", vm.Sso.Methods)
	}
}