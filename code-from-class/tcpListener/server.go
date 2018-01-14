package main

import (
	"context"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

var counter int64

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	stopChan := make(chan struct{})
	go server(ctx, stopChan, l)

	logrus.Info("Server started")
	<-ctx.Done()
	logrus.Info("Timeout")
	l.Close()
	logrus.Info("Listener closed")
	<-stopChan
	logrus.Info("Server stoped")
}

func server(ctx context.Context, stopChan chan struct{}, l net.Listener) {
	defer close(stopChan)
	handlerWG := &sync.WaitGroup{}
	for {
		select {
		case <-ctx.Done():
			logrus.Info("Server stop, waiting for handlers")
			handlerWG.Wait()
			return
		default:
			c, err := l.Accept()
			if err != nil {
				continue
			}
			logrus.Info("New connect! Woohoo!")
			// TODO: HANDLE!
			handlerWG.Add(1)
			go handler(handlerWG, c)
		}
	}
}

func handler(handlerWG *sync.WaitGroup, c net.Conn) {
	defer handlerWG.Done()
	counter = atomic.AddInt64(&counter, 1)
	c.Write([]byte("ok" + strconv.Itoa(int(counter))))
	c.Close()
}
