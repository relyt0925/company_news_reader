package job

import (
	"time"
	"sync"
	"github.com/pkg/errors"
	"fmt"
	"github.com/relyt0925/company_news_reader/myallies"
	"github.com/relyt0925/company_news_reader/routinepool"
)

var MapLock sync.RWMutex
var activeCompanies map[string]bool
var ticker *time.Ticker


func resultLogger(outputChan <-chan routinepool.ExecutableResult ){
	for{
		select{
		case result := <-outputChan:
			if res, err := result.Result(); err != nil{
				fmt.Println("IM THE BEST!!!!")
				fmt.Println(err.Error())
			} else{
				fmt.Println(res.(string))
			}
			return
		}
	}
}


func init(){
	MapLock= sync.RWMutex{}
	activeCompanies = make(map[string]bool)
	ticker = time.NewTicker(time.Second * 5)
	go func() {
		submitter := routinepool.NewSubmitter(50, 50)
		for {
			<-ticker.C
			fmt.Println(activeCompanies)
			MapLock.RLock()
			for k, _ := range activeCompanies{
				fmt.Println("COMPANY "+k)
				outChan := make(chan routinepool.ExecutableResult)
				stockJob := myallies.StockQuoteJob{Job:routinepool.Job{OutputQueue:outChan},Company:k}
				submitter.Submit(&stockJob)
				go resultLogger(outChan)
			}
			MapLock.RUnlock()
		}
	}()
}


func AddCompany(company string) error{
	MapLock.Lock()
	defer MapLock.Unlock()
	if _, exists := activeCompanies[company]; exists{
		return errors.New(fmt.Sprint("Company %s already is being monitored", company))
	}
	activeCompanies[company]=true
	fmt.Println(activeCompanies)
	return nil
}


