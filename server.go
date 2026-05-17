package main

import (
	_ "encoding/json"
	_ "fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	MAX_CLIENTS = 100

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	readWait  = 10 * time.Minute

	pongWait   = 45 * time.Second
	pingPeriod = 40 * time.Second

	maxMessageSize = 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//---------------------------------------------------------------------------------------------------------------------------------------------

func handleHttp(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		log.Println("GET request received")

		keys := r.URL.Query()
		id := keys.Get("id")

		log.Printf("GET id: %s", id)

		//session := keys.Get("session")

		//TODO проверка id, session

		/*client := clients.getClient(id) //ждущий клиент
		if client == nil {
			//fmt.Fprintf(w, "end")
			return
		}

		if !client.session {
			log.Println("Client found; start session")

			fmt.Fprintf(w, "session")
			//сообщить клиенту о начале сессии
			client.send <- createSessionStartMessage()
			client.session = true
		}

		fmt.Fprintf(w, client.command)*/

	case "POST":
		log.Println("POST request received")

		if err := r.ParseForm(); err != nil {
			log.Printf("ParseForm() err: %v", err)
			return
		}

		id := r.Form.Get("id")
		r.Form.Del("id")

		log.Printf("POST id: %s", id)

		//TODO проверка ID

		/*client := clients.getClient(id) //ждущий клиент
		if client == nil {
			return
		}

		var storage StorageAE2
		var data string
		for data, _ = range r.Form {
			//log.Println(data)

			err := json.Unmarshal([]byte(data), &storage)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}

		//log.Println(storage.Items[1].Name)
		//log.Println(data)

		//TODO проверка ID

		if id == storage.ID {
			client.send <- data
		}*/

	default:
	}
}
