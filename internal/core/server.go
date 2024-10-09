package core

import (
	"log"
	"net/http"
	"tic-tac-toe/internal/common"
	"tic-tac-toe/internal/game"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	clients    map[*common.Client]bool
	register   chan *common.Client
	unregister chan *common.Client
	receive    chan ClientMessage
}

func (server *Server) Start(addr string) {
	go server.startWorker()

	http.HandleFunc("/ws", server.onNewConnection)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (server *Server) onNewConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := common.NewClient(conn)

	onUnregister := func() {
		server.unregister <- client
	}

	onReceive := func(message []byte) {
		server.receive <- ClientMessage{
			Content: message,
			Client:  client,
		}
	}

	client.Bind(onUnregister, onReceive)

	server.register <- client
}

func (server *Server) startWorker() {
	module := game.NewGameModule()
	connection := game.NewGameConnectionManager(module)

	for {
		select {
		case client := <-server.register:
			server.clients[client] = true
			connection.OnNewConnection(client)
		case client := <-server.unregister:
			connection.OnDisconnected(client)
			delete(server.clients, client)
		case message := <-server.receive:
			connection.OnMessage(message.Client, message.Content)
		}
	}
}

func NewServer() *Server {
	return &Server{
		register:   make(chan *common.Client),
		unregister: make(chan *common.Client),
		clients:    make(map[*common.Client]bool),
		receive:    make(chan ClientMessage),
	}
}

type ClientMessage struct {
	Content []byte
	Client  *common.Client
}
