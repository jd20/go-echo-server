package main

import (
	"net/http"
	_ "encoding/json"
	"flag"

	log "github.com/Sirupsen/logrus"
	_ "github.com/codegangsta/negroni"
	"github.com/gorilla/websocket"
)

var parseJSON bool

var upgrader = websocket.Upgrader{
	//ReadBufferSize:  1024,
	//WriteBufferSize: 1024,
}

type Message struct {
	Topic   string      `json:"topic"`
	Event   string      `json:"event"`
	Ref     string      `json:"ref"`
	Payload interface{} `json:"payload"`
}

func main() {
	flag.BoolVar(&parseJSON, "parse", false, "parse and encode message payloads as JSON")
	flag.Parse()
	log.Infof("Parsing JSON: %v", parseJSON)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/socket/websocket", handleWebsocket)

	//n := negroni.Classic()
	//n.UseHandler(mux)
	//n.Run(":4000")

	http.HandleFunc("/socket/websocket", handleWebsocket)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	//if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	//}

	/*
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithField("err", err).Println("Upgrading to websockets")
		http.Error(w, "Error Upgrading to websockets", 400)
		return
	}
	ws.WriteMessage(websocket.CloseMessage, []byte{})

	connected := true
	for connected {
		mt, data, err := ws.ReadMessage()
		if err != nil {
			log.Info("Websocket closed!")
			break
		}
		switch mt {
		case websocket.TextMessage, websocket.BinaryMessage:
			// If parsing the message payload, do that now
			if parseJSON {
				var msg Message
				if err := json.Unmarshal(data, &msg); err != nil {
					log.Error("Error parsing JSON message")
					break
				}
				if data, err = json.Marshal(msg); err != nil {
					log.Error("Error encoding JSON message")
					break
				}
			}

			// Echo the data back to the user
			if err = ws.WriteMessage(mt, data); err != nil {
				log.Error("Error writing websocket message")
			}
		default:
			log.Warning("Unknown Message!")
		}
	}
	ws.WriteMessage(websocket.CloseMessage, []byte{})
	*/
}
