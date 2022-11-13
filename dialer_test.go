package IDClient

import "testing"

func TestDialer_DialWithContext(t *testing.T) {
	dialer := &Dialer{
		MaxPoolSize: 50000,
		MinPoolSize: 20000,
		DcID:        0,
		WorkerId:    0,
		Addr:        "127.0.0.1:11451",
	}
	client, err := dialer.Dial()
	if err != nil {
		t.Fatal("dial fatal:", err)
	}
	defer client.Close()

	id, err := client.GetID()
	if err != nil {
		t.Fatal(err)
	}

	client.Recycle(id)
}
