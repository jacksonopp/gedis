package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		send(ctx, input)
		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}

func send(ctx context.Context, msg string) {
	done := make(chan string)

	go func() {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatalln("error dialing tcp server", err)
		}
		if ctx.Err() != nil {
			log.Println("Context error:", ctx.Err())
			return
		}
		defer conn.Close()

		data := []byte(msg)
		_, err = conn.Write(data)
		if err != nil {
			log.Println("error:", err)
		}
		if ctx.Err() != nil {
			log.Println("Context error:", ctx.Err())
			return
		}

		// Read and process data from the server
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error:", err)
			return
		}
		if ctx.Err() != nil {
			log.Println("Context error:", ctx.Err())
			return
		}

		done <- string(buffer[:n])
	}()

	select {
	case result := <-done:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("Work cancelled", ctx.Err())
	}
}
