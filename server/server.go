package server

import (
	"github.com/gorilla/websocket"
	"github.com/sintell/wsgame/connection"
	"github.com/sintell/wsgame/utils"
	"log"
	"net/http"
	"os"
)

type Server struct {
	clients map[string]*Client
	Logger  *utils.Logger
	Cfg     *Configurator
}

func New() *Server {
	loggerFlags := log.Ldate | log.Ltime | log.Lmicroseconds
	logger := utils.NewLogger(utils.DEBUG, os.Stdout, loggerFlags)

	return &Server{
		clients: make(map[string]*Client),
		Logger:  logger,
		Cfg:     NewConfigurator("config.json"),
	}
}

func (s *Server) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Scheme != "ws" {
			s.Logger.Error("Wrong scheme. Must be ws, but got %s", r.URL.Scheme)
		}
		s.Logger.Info("New connection from %s", r.Header.Get("Origin"))
		client := connection.New(w, r)
		s.Loop(client)
	}
}

func (s *Server) HandleMessage(message *connection.Message, client *websocket.Conn) {
	switch message.Type {
	case connection.SystemMsg:
		{
		}
	case connection.PingMsg:
		{
		}
	case connection.DataMsg:
		{
			s.Logger.Info("Got data from client %s: %i", client.UnderlyingConn().LocalAddr().String(), message.Data)
		}

	default:
		{
			s.Logger.Debug("This is debug message")
			s.Logger.Warning("Got wrong message type from client: %s", client.UnderlyingConn().LocalAddr().String())
			s.Logger.Error("Can't perform auth on user")
		}
	}
}

func (s *Server) Loop(client *websocket.Conn) {
	message := connection.NewMessage()
	for {
		err := client.ReadJSON(message)
		if err != nil {
			s.Logger.Error("Cant read JSON request from client. Reason: %s", err.Error())
			panic(err.Error())
		}

		s.HandleMessage(message, client)
	}
}
