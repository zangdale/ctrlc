package server

import (
	"context"
	"ctrlc/db"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"

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

	count := QueryInt(req.URL, "count", 20)
	offset := QueryInt(req.URL, "offset", 0)
	searchStr := QueryStr(req.URL, "search", "")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	res, err := db.GetMsgs(count, offset, searchStr)
	if err != nil {
		resp.Info().Err(err).Str("handle", "get").Msg("error")
		return
	}

	resp.Info().Str("handle", "get").Interface("data", res).Msg("ok")
}

func QueryInt(u *url.URL, key string, def int) int {
	countStr := u.Query().Get(key)
	res := def
	if countStr != "" {
		i, _ := strconv.Atoi(countStr)
		if i > 0 {
			res = i
		}
	}
	return res
}
func QueryStr(u *url.URL, key string, def string) string {
	countStr := u.Query().Get(key)
	res := def
	if countStr != "" {
		res = countStr
	}
	return res
}

func getNetIpsHandle(w http.ResponseWriter, req *http.Request) {
	resp := zerolog.New(w)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		resp.Info().Err(err).Str("handle", "getip").Msg("error")
		return
	}
	var ips []string
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	resp.Info().Str("handle", "getip").Strs("data", ips).Msg("ok")
}

func addMsgHandle(w http.ResponseWriter, req *http.Request) {
	resp := zerolog.New(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if req.Method != http.MethodPost {
		resp.Info().Str("err", "method not post").Str("handle", "add").Msg("error")
		return
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		resp.Info().Err(err).Str("handle", "add").Msg("error")
		return
	}
	var msg = struct {
		Msg string `json:"msg"`
	}{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		resp.Info().Err(err).Str("handle", "add").Msg("error")
		return
	}
	AddMsg(msg.Msg)
	err = db.SaveCopyMsg(context.TODO(), msg.Msg)
	if err != nil {
		resp.Info().Err(err).Str("handle", "add").Msg("error")
		return
	}

	resp.Info().Str("handle", "add").Str("data", "ok").Msg("ok")
}
