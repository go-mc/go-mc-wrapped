// noinspection GoMultiplePackages
package main

import (
	"encoding/json"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/net"
	wrapped "github.com/go-mc/go-mc-wrapped"
	"github.com/google/uuid"
	"log"
)

func main() {
	l, err := net.ListenMC(":25565")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Accept error: %v", err)
		}
		go acceptConn(wrapped.Wrap(conn))
	}
}

func acceptConn(conn wrapped.Conn) {
	// noinspection GoUnhandledErrorResult
	defer conn.Close()
	// handshake
	_, intention, err := handshake(conn)
	if err != nil {
		log.Printf("Handshake error: %v", err)
		return
	}

	switch intention {
	default: //unknown error
		log.Printf("Unknown handshake intention: %v", intention)
	case 1: //for status
		acceptListPing(conn)
	}
}

// handshake receive and parse Handshake packet
func handshake(conn wrapped.Conn) (int, int, error) {
	packet := wrapped.HandshakingSetProtocolToServer{}
	err := conn.Receive(&packet)
	return packet.ProtocolVersion, packet.NextState, err
}

func acceptListPing(conn wrapped.Conn) {
	for i := 0; i < 2; i++ { // ping or list. Only accept twice
		p, err := conn.ReadPacket()
		if err != nil {
			return
		}

		switch p.ID {
		case 0x00: //List
			err = conn.Send(wrapped.StatusServerInfoToClient{Response: listResp()})
		case 0x01: //Ping
			err = conn.WritePacket(p)
		}
		if err != nil {
			return
		}
	}
}

type player struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

// listResp return server status as JSON string
func listResp() string {
	var list struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`
		Players struct {
			Max    int      `json:"max"`
			Online int      `json:"online"`
			Sample []player `json:"sample"`
		} `json:"players"`
		Description chat.Message `json:"description"`
		FavIcon     string       `json:"favicon,omitempty"`
	}

	list.Version.Name = "Chat Server"
	list.Version.Protocol = 736
	list.Players.Max = 20
	list.Players.Online = 123
	list.Players.Sample = []player{} // must init. can't be nil
	list.Description = chat.Message{Text: "Powered by go-mc", Color: "blue"}

	data, err := json.Marshal(list)
	if err != nil {
		log.Panic("Marshal JSON for status checking fail")
	}
	return string(data)
}
