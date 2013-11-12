package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "go home, you're not a socket, you're drunk", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	log.Println("new client")
	defer log.Println("client closed")

	conn.WriteMessage(websocket.TextMessage, []byte("hello"))

	// TODO: handle timing out; don't want to have these linger
	for {
		typ, msg, err := conn.ReadMessage()
		log.Println(string(msg))

		if err != nil {
			return
		}
		if err := conn.WriteMessage(typ, msg); err != nil {
			return
		}
	}
}
