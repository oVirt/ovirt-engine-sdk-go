package ovirtsdk4

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		t.Fatalf("Make connection failed, reason: %s", err.Error())
	}
	defer conn.Close()
	clustersListResponse, err2 := conn.SystemService().ClustersService().
		List().
		CaseSensitive(false).
		Max(100).
		Send()

	if err2 != nil {
		t.Fatalf("Get clusters failed, reason: %s", err2.Error())
	}

	for _, cluster := range clustersListResponse.Clusters() {
		t.Logf("cluster(%v): CPU architecture is %v and type is %v", *cluster.Id, cluster.Cpu.Architecture, *cluster.Cpu.Type)
	}

}
