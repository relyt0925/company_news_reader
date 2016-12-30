package main

import (
	"github.com/relyt0925/company_news_reader/job"
	"time"
	"fmt"
)

/*
func resultLogger(outputChan <-chan routinepool.ExecutableResult ){
	for{
		select{
		case result := <-outputChan:
			fmt.Println("IM THE BEST!!!!")
			if res, err := result.Result(); err != nil{
				fmt.Println(err.Error())
			} else{
				fmt.Println(res.(string))
			}
			return
		}
	}
}


func main(){
	submitter := routinepool.NewSubmitter(50,20)
	outChan := make(chan routinepool.ExecutableResult)
	job := myallies.StockQuoteJob{Job:routinepool.Job{OutputQueue:outChan},Company:"MSFT"}
	go resultLogger(outChan)
	submitter.Submit(&job)
	time.Sleep(time.Second * 8)
	submitter.Cleanup()
	time.Sleep(time.Second * 2)
	return
}
*/

func main(){
	err := job.AddCompany("MSFT")
	if err != nil{
		fmt.Println(err)
	}
	job.AddCompany("AAPL")
	job.AddCompany("GOOG")
	time.Sleep(time.Second * 300)
}