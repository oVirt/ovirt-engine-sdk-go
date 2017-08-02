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

	// Get the reference to the "datacenters" service:
	datacentersService := conn.SystemService().DataCentersService()

	// Use the "list" method of the "datacenters" service to list all the datacenters of the system:
	datacentersResponse, err := datacentersService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get datacenter list, reason: %v\n", err)
		return
	}

	// Print the datacenter names and identifiers:
	for _, dc := range datacentersResponse.DataCenters() {
		fmt.Printf("Datacenter - (name: %v, id: %v)\n", *dc.Name, *dc.Id)
		fmt.Printf("  Supported versions are: ")
		for _, sv := range dc.SupportedVersions {
			fmt.Printf("(Major: %v, Minor: %v)  ", *sv.Major, *sv.Minor)
		}
		fmt.Println("")
	}

}
