package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeoutFlag string

func init() {
	flag.StringVar(&timeoutFlag, "timeout", "10s", "connection timeout")
}

func main() {
	flag.Parse()

	host := flag.Args()[0]
	port := flag.Args()[1]
	if host == "" || port == "" {
		log.Fatal("host or port were not defined")
	}

	timeout, err := time.ParseDuration(timeoutFlag)
	if err != nil {
		log.Fatal("timeout was not parsed correctly")
	}

	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := client.Receive(); err != nil {
			cancel()
		}
	}()

	go func() {
		if err := client.Send(); err != nil {
			cancel()
		}
	}()

	select {
	case <-sigCh:
		cancel()
		signal.Stop(sigCh)
		return
	case <-ctx.Done():
		return
	}
}
