package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type message struct {
	Handle string `json:"handle"`
	Text   string `json:"text"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/socket/websocket", handleWebsocket)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithField("err", err).Println("Upgrading to websockets")
		http.Error(w, "Error Upgrading to websockets", 400)
		return
	}

	connected := true
	for connected {
		mt, data, err := ws.ReadMessage()
		if err != nil {
			log.Info("Websocket closed!")
			break
		}
		switch mt {
		case websocket.TextMessage:
			log.Info("Received text: " + string(data))
			if err = ws.WriteMessage(mt, data); err != nil {
				log.Error("Error writing websocket message")
			}
		case websocket.BinaryMessage:
			log.Info("Received binary: " + string(data))
			if err = ws.WriteMessage(mt, data); err != nil {
				log.Error("Error writing websocket message")
			}
		default:
			log.Warning("Unknown Message!")
		}
	}
	ws.WriteMessage(websocket.CloseMessage, []byte{})
}
