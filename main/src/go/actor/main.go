package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type ping struct {
}

type pong struct {
}

type pingActor struct {
	pongPid *actor.PID
}

func (p *pingActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *ping:
		fmt.Println("pong!")
	case struct{}:
		fmt.Println("ping Received, Send pong message")
		ctx.Send(p.pongPid, &pong{})
	default:
		fmt.Println("msg", msg)
	}
}

func main() {
	sys := actor.NewActorSystem()
	// pong
	pongRecv := func(ctx actor.Context) {
		switch ctx.Message().(type) {
		case *pong:
			fmt.Println("pong Received, Send ping message")
		}
	}
	pongProps := actor.PropsFromFunc(pongRecv)
	pongPid, err := sys.Root.SpawnNamed(pongProps, "PONG")
	if err != nil {
		panic(err)
	}

	// ping
	pingProps := actor.PropsFromProducer(func() actor.Actor {
		return &pingActor{
			pongPid: pongPid,
		}
	})
	pingPid := sys.Root.Spawn(pingProps)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			sys.Root.Send(pingPid, struct{}{})
			// sys.Root.Send(pingPid, &ping{})
		case <-ch:
			log.Print("Finish")
			return

		}
	}
}
