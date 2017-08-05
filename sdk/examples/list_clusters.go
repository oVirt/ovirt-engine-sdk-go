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

	// Get the reference to the "clusters" service:
	clustersService := conn.SystemService().ClustersService()

	// Use the "list" method of the "clusters" service to list all the clusters of the system:
	clustersResponse, err := clustersService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get cluster list, reason: %v\n", err)
		return
	}

	// Print the datacenter names and identifiers:
	for _, clu := range clustersResponse.Clusters() {
		fmt.Printf("Cluster - (name: %v, id: %v)\n", *clu.Name, *clu.Id)
		// fmt.Printf("Cluster is %+v\n", clu)
	}

}
