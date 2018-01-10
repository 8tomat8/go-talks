package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64

func handler(c net.Conn) {
	counter = atomic.AddInt64(&counter, 1)
	c.Write([]byte("ok" + strconv.Itoa(int(counter)) + "\n"))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	stopChan := make(chan struct{})
	go server(ctx, stopChan, l)

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			client()
			wg.Done()
		}()
	}

	wg.Wait()
	l.Close()
	cancel()
	<-stopChan
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
		return
	}

	data := make([]byte, 2<<10)
	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Cli error:", err)
			break
		}
		fmt.Println(string(data[:n]))
	}
}

func server(ctx context.Context, stopChan chan struct{}, l net.Listener) {
	defer func() { close(stopChan) }()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Server stop")
			return
		default:
			c, err := l.Accept()
			if err != nil {
				continue
			}
			// TODO: HANDLE!
			go handler(c)
		}
	}
}
