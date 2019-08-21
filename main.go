package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/calvernaz/loraserver-debug/net/udp"
)

func main() {
	// the exit channel - blocks the main thread
	exitCh := make(chan struct{})


	// build udp server
	udpServer := &udp.UDPServer{
		1700,
	}

	// read messages<
	for msg := range udpServer.ReadUDPMessages() {
		fmt.Println(msg)
	}

	// install ctrl-c or docker stop
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-signalCh:
			exitCh <- struct{}{}
			return
		}
	}()

	<-exitCh
	logrus.Info("the work is done here, exiting...")
}
