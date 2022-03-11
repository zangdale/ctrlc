package server

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

var msgsChans = make(map[chan string]bool)

func AddMsg(txt string) {
	for m := range msgsChans {
		m <- txt
	}
}

func getChan() chan string {
	msgs := make(chan string)
	msgsChans[msgs] = true
	return msgs
}

func delChan(ch chan string) {
	delete(msgsChans, ch)
}

// https://developer.mozilla.org/zh-CN/docs/Web/API/EventSource
// 结构信息 https://javascript.ruanyifeng.com/htmlapi/eventsource.html

func msgHandle(w http.ResponseWriter, req *http.Request) {
	// resp := zerolog.New(w)

	ctx := req.Context()
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")

	msgs := getChan()
	defer delChan(msgs)

	for {
		select {
		case msg := <-msgs:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			f.Flush()
		case <-ctx.Done():
			return
		}
	}

}

func getHandle(w http.ResponseWriter, req *http.Request) {
	resp := zerolog.New(w)
	resp.Info().Str("handle", "get").Msg("ok")
}

func searchHandle(w http.ResponseWriter, req *http.Request) {
	resp := zerolog.New(w)
	resp.Info().Str("handle", "search").Msg("ok")
}
