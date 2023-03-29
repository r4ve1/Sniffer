package reader

import (
	"testing"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/stretchr/testify/assert"
)

type testV struct {
}

func Test_Efficiency(t *testing.T) {
	h, err := pcap.OpenOffline("D:/Desktop/large.pcap")
	assert.Nil(t, err)
	src := gopacket.NewPacketSource(h, h.LinkType())
	cnt := 0
	start := time.Now()
	for pkg := range src.Packets() {
		cnt++
		_ = pkg
	}
	t.Logf("cnt: %d", cnt)
	t.Logf("time: %s", time.Since(start))
}
