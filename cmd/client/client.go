package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go dial(ctx, "first", &wg)
	go dial(ctx, "second", &wg)
	wg.Wait()
}

func dial(ctx context.Context, msg string, wg *sync.WaitGroup) {
	defer wg.Done()

	if ctx.Err() == context.DeadlineExceeded {
		log.Panic("DEADLINE EXCEEDED")
	}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("error dialing tcp server", err)
	}
	defer conn.Close()

	data := []byte(msg)
	_, err = conn.Write(data)
	if err != nil {
		log.Println("error:", err)
	}

	// Read and process data from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Received:", string(buffer[:n]))
}
