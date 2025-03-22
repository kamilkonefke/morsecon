package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func ws_handler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }

        if err := conn.WriteMessage(websocket.TextMessage,  msg); err != nil {
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", ws_handler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("ListenAndServe: ", err)
        return
    }
}
