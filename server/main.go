package main

import (
	"log"
	"net/http"

	"github.com/Dejannnn/websocketapp.git/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket endpoint start...")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println("Error: ", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool}

	pool.Register <- client
	client.Read()
}
func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		serveWs(pool, w, r)

	})
}
func main() {
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
