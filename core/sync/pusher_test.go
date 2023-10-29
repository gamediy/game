package sync

import (
	"game/utility/utils/xpusher"
	"testing"
)

func TestMessage_Trigger(t *testing.T) {
	xpusher.Init("1696531", "bb8063bd4d4939b6c4ea", "366111b02e0482f92d16", "mt1")
	msg := Message{}
	if err := msg.Trigger(ChannelAdmin, EventDeposit); err != nil {
		t.Fatal(err)
	}
}
