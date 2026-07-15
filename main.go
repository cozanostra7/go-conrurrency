package main

import (
	"context"
	"fmt"
	"go_async/mailman"
	"go_async/miner"
	"time"
)

func main() {

	var coal int
	var mails []string

	minerContext,minerCancel:= context.WithCancel(context.Background())
	postmaniContext,postmanCancel :=context.WithCancel(context.Background())

	go func () {

		time.Sleep(3*time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
	}()
	
	coalTransferPoint:= miner.MinerPool(minerContext,2)
	

	mailTransferPoint:= mailman.MailmanPool(postmaniContext,2)

	isCoalClosed:=false
	isMailClose:=false
	for !isCoalClosed || !isMailClose{

		select {
		case c,ok:=<-coalTransferPoint:
			if!ok {
				isCoalClosed = true
				continue

			}
		coal+=c
		case m,ok:=<-mailTransferPoint:
			if !ok {
			isMailClose =true
			continue
			}
			mails = append(mails,m)
	}
	 
	}
	
	fmt.Println(" Total coal: ",coal)
	fmt.Println("total mails: ",mails)

}