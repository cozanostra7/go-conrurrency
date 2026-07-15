package mailman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Mailman(ctx context.Context,wg *sync.WaitGroup,trasferPoint chan<- string,n int,mail string){
	defer wg.Done()
	
	for {
		select {
		case<-ctx.Done():
			fmt.Println("mailman n: ",n,"the worker has finished")
			return
		default:
			fmt.Println("Im mailman , i delivered psot: ",n)
			time.Sleep(1*time.Second)
			fmt.Println("Im mailman n: ",n,"delivered : ",mail)

			trasferPoint<-mail

			fmt.Println("Im mailman : ",n,"delivered: ",mail)
		}

	}
}


func MailmanPool(ctx context.Context,mailmanCoint int) <-chan string {
	mailTransferPoint:= make(chan string)

	wg:=&sync.WaitGroup{}

	for i:=1;i <= mailmanCoint; i++ {
		wg.Add(1)

		go Mailman(ctx,wg,mailTransferPoint,i,postmantoMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmantoMail(postmanNumber int) string {

	pstm:=map[int] string {
		1:"Job",
		2:"Payment",
		3:"parents",
	}

	mail,ok:=pstm[postmanNumber]
	if!ok {
		return "lottery"
	}
	return mail
}