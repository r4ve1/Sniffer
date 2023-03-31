package reader

import (
	"bufio"
	"io"
	"log"
	"net/http"

	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"

	"sniffer/internal/view"
)

// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct {
	v         *view.View
	isHttpReq func(gopacket.Flow) bool
}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
	v              *view.View
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	stream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
		v:         h.v,
	}

	if h.isHttpReq(stream.transport) {
		go stream.runRequest() // Important... we must guarantee that data from the reader stream is read.
	} else {
		go stream.runResponse() // Important... we must guarantee that data from the reader stream is read.
	}

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &stream.r
}

func (h *httpStream) runResponse() {

	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		resp, err := http.ReadResponse(buf, nil)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
			return
		} else {
			bodyLen := tcpreader.DiscardBytesToEOF(resp.Body)
			log.Println("Body contains", bodyLen, "bytes")
			h.v.RenderHttpRespBrief(h.net, resp)
			//_ = resp.Body.Close()
			//printResponse(resp, h, bodyBytes)
			// log.Println("Received response from stream", h.net, h.transport, ":", resp, "with", bodyBytes, "bytes in response body")
		}
	}
}
func (h *httpStream) runRequest() {
	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			//h.v.RenderHttpReqBrief(h.transport, req)
			h.v.RenderHttpReqBrief(h.net, req)
			_ = req.Body.Close()
			_ = tcpreader.DiscardBytesToEOF(req.Body)
			//printRequest(req, h, bodyBytes)
			// log.Println("Received request from stream", h.net, h.transport, ":", req, "with", bodyBytes, "bytes in request body")
		}
	}
}
