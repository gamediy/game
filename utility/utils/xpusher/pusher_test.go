package xpusher

import (
	"testing"
)

func init() {
	Init("1696513", "789cc87800cc45174098", "021f122d02bf1feada17", "ap3")
}
func TestTrigger(t *testing.T) {
	data := map[string]string{"message": "hello world"}
	err := Trigger("my-channel", "my-event", data)
	if err != nil {
		t.Fatal(err)
	}
}
