package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferpoint chan<-int,
	n int,
	power int) {
	
	for {
		select{
		case <-ctx.Done():
			fmt.Println("Worker: ",n,"work has finished")
			return
		default:
			fmt.Println("Miner n: ",n)
			time.Sleep(1*time.Second)
			fmt.Println("Miner n: ","i got some mine")
			transferpoint<-power
			fmt.Println("Miner n: ",n,"I transfered : ",power)

			}
		
	}
}


func MinerPool(ctx context.Context,minerCount int) <-chan int{
	coalTransferPoint:=make(chan int)

	wg:=&sync.WaitGroup{}

	for i:=1; i<=minerCount;i++{
		wg.Add(1)
		go Miner(ctx,wg,coalTransferPoint,i,i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()
	
	return coalTransferPoint
}