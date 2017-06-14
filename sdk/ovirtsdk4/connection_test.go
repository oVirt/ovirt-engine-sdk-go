package ovirtsdk4

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"
	conn, err := NewConnection(
		inputRawURL, "admin@internal", "qwer1234",
		"", true, "", false,
		uint64(10*time.Second), true)
	if err != nil {
		t.Errorf("connection failed, reason %s", err.Error())
	}
	// ovRequest := NewOvRequest("GET", "/clusters", nil, nil, "")
	clusterList, err := conn.SystemService().ClustersService().List(false, false, 100, "", nil, nil, false)
	for _, cluster := range clusterList {
		t.Logf("cluster(%v): CPU architecture is %v and type is %v", cluster.Id, cluster.Cpu.Architecture, cluster.Cpu.Type)
	}

	// ovResponse, err := conn.Send(ovRequest)
	// t.Logf("response %s", string(ovResponse.Body))
}
