package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
    "os/exec"
)

func main(){
    /*
    out, err := exec.Command("/bin/python3", "myPythonFile.py", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - -").Output()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("we got an output: ", out)
    fmt.Println("ascii output to string: ", string([]byte(out)))
	*/
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
            recievedPosition(ws, splitStrings[1])
        } else if pString == "gimmeNewPosition" {
            sendNewPosition(ws, "8/8/8/8/4P3/8/8/8 b - - 0 1")
        } else {
            fmt.Println("shit dammit missed something")
        }
    }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func recievedPosition(ws *websocket.Conn, fenString string) {
    fmt.Println("recieved chess position: ", fenString)
    out, err := exec.Command("/bin/python3", "myPythonFile.py", fenString).Output()
    if err != nil {
        fmt.Println("shit fucked up when calculating new position")
        fmt.Println(err)
        return
    }
    fenAfterCalculation := string([]byte(out))
    sendNewPosition(ws, fenAfterCalculation)
}

func sendNewPosition(ws *websocket.Conn, fenString string) {
    err := ws.WriteMessage(1, []byte("updatePosition " + fenString))
    if err != nil {
        fmt.Println(err)
        return
    }
}