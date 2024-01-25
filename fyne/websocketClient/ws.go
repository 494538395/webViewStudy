package main

import (
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn

func connectTOWs(addr string, headers map[string]string, param map[string]string) (*websocket.Conn, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	// params
	params := url.Values{}
	for k, v := range param {
		params.Set(k, v)
	}
	u.RawQuery = params.Encode()

	// headers
	header := http.Header{}

	for k, v := range headers {
		header.Add(k, v)
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return nil, err
	}

	return c, nil
}
