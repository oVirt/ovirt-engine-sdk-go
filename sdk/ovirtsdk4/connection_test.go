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
	clusterList, err := conn.SystemService().ClustersService().List(false, false, 100, "", nil, nil, false)
	if err != nil {
		t.Fatalf("Get clusters failed, reason: %s", err.Error())
	}
	for _, cluster := range clusterList {
		t.Logf("cluster(%v): CPU architecture is %v and type is %v", *cluster.Id, cluster.Cpu.Architecture, *cluster.Cpu.Type)
	}
}
