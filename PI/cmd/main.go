package main

import (
	pi "PI/internal"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	//gorutinsAmount := pflag.IntP("gorutins_amount", "n", 1, "amount of gorutins")
	gorutinsAmount := 21
	result := make(chan float64)
	quit := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < gorutinsAmount; i++ {
		go pi.LeibnicSum(result, quit, i, gorutinsAmount)
	}
	pi := 0.0
	go func() {
		for i := 0; i < gorutinsAmount; i++ {
			pi += <-result
		}
		fmt.Print("\n", pi)
		wg.Done()
	}()
	<-sigs
	for i := 0; i < gorutinsAmount; i++ {
		quit <- true
	}
	wg.Wait()
}
