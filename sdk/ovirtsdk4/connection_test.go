package ovirtsdk4

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	inputRawUrl := "https://10.1.111.229/ovirt-engine/api"
	conn, err := NewConnection(
		inputRawUrl, "admin@internal", "qwer1234",
		nil, true, nil, false,
		10 * time.Second, true)
	if err != nil {
		t.Errorf("connection failed, reason %s", err.Error())
	}
	ovRequest := NewOvRequest("GET", "/clusters", nil, nil, nil)
	ovResponse, err := conn.Send(ovRequest)
	t.Logf("response %s", string(ovResponse.Body))
}