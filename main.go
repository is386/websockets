package main

import (
	"fmt"
	"net/http"

	"github.com/is386/super-base-64/superbase64"
)

const (
	hostname = "localhost"
	port     = 8080
	uri      = "/"
)

func handleUpgrade(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "method of request MUST be GET", http.StatusBadRequest)
		return
	}

	if !req.ProtoAtLeast(1, 1) {
		http.Error(w, "HTTP version MUST be at least 1.1", http.StatusBadRequest)
		return
	}

	if req.RequestURI != uri {
		http.Error(w, "resource uri does not match", http.StatusBadRequest)
		return
	}

	if req.Host != fmt.Sprintf("%s:%d", hostname, port) {
		http.Error(w, "Host header missing or does not match hostname", http.StatusBadRequest)
		return
	}

	if req.Header.Get("Upgrade") != "websocket" {
		http.Error(w, "Upgrade header missing or does not equal 'websocket'", http.StatusBadRequest)
		return
	}

	if req.Header.Get("Connection") != "Upgrade" {
		http.Error(w, "Connection header missing or does not equal 'Upgrade'", http.StatusBadRequest)
		return
	}

	decoded, err := superbase64.NewStdEncoding().Decode(req.Header.Get("Sec-WebSocket-Key"))
	if err != nil || len(decoded) != 16 {
		http.Error(w,
			"Sec-WebSocket-Key header missing or is not a 16 byte base64 encoded string",
			http.StatusBadRequest,
		)
		return
	}

	// TODO: Figure out what exactly I need to verify here
	if req.Header.Get("Origin") != "" {
		fmt.Println("Origin included")
	}

	if req.Header.Get("Sec-WebSocket-Version") != "13" {
		http.Error(w, "Sec-WebSocket-Version header missing or is not '13'", http.StatusBadRequest)
		return
	}

	// TODO: Figure out what exactly I need to verify here
	if req.Header.Get("Sec-WebSocket-Protocol") != "" {
		fmt.Println("Sec-WebSocket-Protocol included")
	}

	// TODO: Figure out what exactly I need to verify here
	if req.Header.Get("Sec-WebSocket-Extensions") != "" {
		fmt.Println("Sec-WebSocket-Extensions included")
	}

	fmt.Println("valid request!")
}

func main() {
	http.HandleFunc(uri, handleUpgrade)
	http.ListenAndServe(fmt.Sprintf("%s:%d", hostname, port), nil)
}
