package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/jacksonopp/gedis/internal/handlers"
)

type GedisServer struct {
	url     string
	port    string
	handler *handlers.GedisHandler
}

func NewGedisServer(port string) *GedisServer {
	return &GedisServer{
		url:     "localhost",
		port:    port,
		handler: handlers.NewGedisHandler(),
	}
}

func (s *GedisServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.url, s.port))
	if err != nil {
		return err
	}

	fmt.Printf("Gedis server started on %s:%s\n", s.url, s.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.handle(conn)
	}
}

func (s GedisServer) handle(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println("Error:", err)
			}
			return
		}

		input := string(buffer[:n])
		response, err := s.handler.HandleCommand(input)

		if err != nil {
			log.Println("Error:", err)
			conn.Write([]byte(err.Error()))
			return
		}

		conn.Write([]byte(response))
	}
}
