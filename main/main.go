package main

import (
	"github.com/relyt0925/company_news_reader/routinepool"
	"fmt"
	"time"
)

func resultLogger(outputChan <-chan routinepool.ExecutableResult, quitChan <-chan bool ){
	for{
		select{
		case result := <-outputChan:
			fmt.Println("IM THE BEST!!!!")
			fmt.Println(result)
		case <-quitChan:
			return
		}
	}
}

func main(){
	submitter := routinepool.NewSubmitter(50,20,50,resultLogger)
	job := routinepool.StockQuoteJob{Job:&routinepool.Job{},Company:"MSFT"}
	submitter.Submit(job)
	time.Sleep(time.Second * 8)
	return
}