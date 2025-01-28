package main

import (
	"log/slog"
	"net"
)

type Server struct {
	Conf
	peer        map[*Peer]bool
	ln          net.Listener
	addPeerchan chan *Peer
	quitchan    chan struct{}
}
type Conf struct {
	AddrListen string
}

func ServerInit(cfg Conf) *Server {
	if len(cfg.AddrListen) == 0 {
		cfg.AddrListen = ":6969"
	}
	return &Server{
		Conf:        cfg,
		peer:        make(map[*Peer]bool),
		addPeerchan: make(chan *Peer),
		quitchan:    make(chan struct{}),
	}
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", s.AddrListen)
	if err != nil {
		return err
	}
	s.ln = ln

	go s.peersLoop()

	slog.Info("Listening for connections", "address", s.AddrListen)

	return s.acceptConnections()
}

func (s *Server) acceptConnections() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("Error accepting connection", "error", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) peersLoop() {
	for {
		select {

		case <-s.quitchan:
			return
		case p := <-s.addPeerchan:
			s.peer[p] = true
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	peer := NewPeer(conn)
	s.addPeerchan <- peer
	slog.Info("New peer connection", "address", conn.RemoteAddr())
	peer.readLoop()
}

func main() {
	server := ServerInit(Conf{AddrListen: ":3333"})

	server.Listen()
}
