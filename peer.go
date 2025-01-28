package main

import "net"

type Peer struct {
	Conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		Conn: conn,
	}
}

func (p *Peer) readLoop() {
	for {
		buf := make([]byte, 1024)
		_, err := p.Conn.Read(buf)
		if err != nil {
			return
		}
	}
}
