package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func main(){
	
	fmt.Println("main")

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	mux.HandleFunc("/ws", wsEndpoint)
	server.ListenAndServe()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    log.Println("Client Connected")
    err = ws.WriteMessage(1, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }
    
    for {
        _, p, err := ws.ReadMessage()
        if err != nil {
            fmt.Println(err)
            return
        }

        pString := string(p)
        fmt.Println(pString)
        splitStrings := strings.Split(pString, " ")


        if splitStrings[0] == "userSentNewPosition" {
            recievedPosition(splitStrings[1])
        } else if pString == "gimmeNewPosition" {
            sendNewPosition(ws)
        } else {
            fmt.Println("shit dammit missed something")
        }
    }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func recievedPosition(fenString string) {
    fmt.Println("recieved chess position: ", fenString)
}

func sendNewPosition(ws *websocket.Conn) {
    err := ws.WriteMessage(1, []byte("updatePosition 8/8/8/8/R7/8/8/8"))
    if err != nil {
        fmt.Println(err)
        return
    }
}