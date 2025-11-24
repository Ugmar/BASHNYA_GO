package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func ChanelWrtie(c chan int, ctx context.Context){
	defer close(c)

	for {
		select{
		case <- ctx.Done():
			return
		case c <- 5:
		}
	}
}

func ChanelRead(c chan int, ctx context.Context, wg *sync.WaitGroup){
	defer wg.Done()

	for{
		select {
		case <- ctx.Done():
			return
		case v := <- c:
			fmt.Println(v)
		}
	}
}

func ChanelWork(n int){
	c := make(chan int)
	var wg sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go ChanelWrtie(c, ctx)

	for range n{
		wg.Add(1)
		go ChanelRead(c, ctx, &wg)
	}

	wg.Wait()
}
