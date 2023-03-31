package view

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	"sniffer/internal/view/renderer"
)

func (it *View) newBrief(timestamp time.Time) *renderer.Brief {
	it.cnt++

	brief := &renderer.Brief{}
	brief.No = it.cnt

	if brief.No == 1 {
		brief.Timestamp = 0
		it.firstPktTimestamp = timestamp.UnixNano()
	} else {
		brief.Timestamp = timestamp.UnixNano() - it.firstPktTimestamp
	}
	return brief
}

func (it *View) packet2Brief(packet gopacket.Packet) *renderer.Brief {
	brief := it.newBrief(packet.Metadata().Timestamp)
	brief.Length = packet.Metadata().Length

	// determine flow
	var flow gopacket.Flow
	if packet.NetworkLayer() != nil {
		flow = packet.NetworkLayer().NetworkFlow()
	} else if packet.LinkLayer() != nil {
		flow = packet.LinkLayer().LinkFlow()
	} else {
		it.log.Error(fmt.Sprintf("No flow found, packet: %v", packet))
	}
	brief.Source = flow.Src().String()
	brief.Destination = flow.Dst().String()
	// determine protocol
	for i := len(packet.Layers()) - 1; i >= 0; i-- {
		l := packet.Layers()[i]
		t := l.LayerType()
		if t != gopacket.LayerTypePayload {
			brief.Protocol = t.String()
			brief.Info = layerInfo(l)
			break
		}
	}

	return brief
}

func (it *View) httpReq2Brief(flow gopacket.Flow, req *http.Request) *renderer.Brief {
	brief := &renderer.Brief{}
	brief.No = it.cnt
	brief.Timestamp = it.last.Metadata().Timestamp.UnixNano() - it.firstPktTimestamp
	brief.Length = it.last.Metadata().Length
	brief.Source = flow.Src().String()
	brief.Destination = flow.Dst().String()
	brief.Protocol = "HTTP"
	brief.Info = fmt.Sprintf("Request: %s %s", req.Method, req.URL)
	brief.Phony = true
	return brief
}

func (it *View) httpResp2Brief(flow gopacket.Flow, resp *http.Response) *renderer.Brief {
	brief := &renderer.Brief{}
	brief.No = it.cnt
	brief.Timestamp = it.last.Metadata().Timestamp.UnixNano() - it.firstPktTimestamp
	brief.Length = it.last.Metadata().Length
	brief.Source = flow.Src().String()
	brief.Destination = flow.Dst().String()
	brief.Protocol = "HTTP"
	brief.Info = fmt.Sprintf("Response: %s", resp.Status)
	brief.Phony = true
	return brief
}

func layerInfo(l gopacket.Layer) string {
	switch l.(type) {
	case *layers.TCP:
		tcp := l.(*layers.TCP)
		var flags []string
		if tcp.FIN {
			flags = append(flags, "FIN")
		}
		if tcp.SYN {
			flags = append(flags, "SYN")
		}
		if tcp.PSH {
			flags = append(flags, "PSH")
		}
		if tcp.ACK {
			flags = append(flags, "ACK")
		}
		f := strings.Join(flags, ",")
		return fmt.Sprintf(
			"%d -> %d [%s] Seq: %d, Ack: %d, Win: %d", tcp.SrcPort, tcp.DstPort, f, tcp.Seq, tcp.Ack, tcp.Window,
		)
	case *layers.UDP:
		udp := l.(*layers.UDP)
		return fmt.Sprintf("%d -> %d", udp.SrcPort, udp.DstPort)
	case *layers.ICMPv4:
		icmp := l.(*layers.ICMPv4)
		return fmt.Sprintf("Type: %d, Code: %d", icmp.TypeCode.Type(), icmp.TypeCode.Code())
	case *layers.ICMPv6:
		icmp := l.(*layers.ICMPv6)
		return fmt.Sprintf("Type: %d, Code: %d", icmp.TypeCode.Type(), icmp.TypeCode.Code())
	default:
		return ""
	}

}
