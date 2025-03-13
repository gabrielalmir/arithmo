package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabrielalmir/arithmo/arithmo"
	"github.com/gabrielalmir/arithmo/router"
)

func main() {
	storage := &arithmo.Storage{}
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nShutting down server...")
		listener.Close()
		os.Exit(0)
	}()

	fmt.Println("Arithmo running on port 6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go router.HandleConnection(conn, storage)
	}
}
