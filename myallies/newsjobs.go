package myallies

import (
	"time"
	"github.com/Shopify/sarama"
	"github.com/relyt0925/company_news_reader/routinepool"
	"github.com/relyt0925/company_news_reader/kafkaproducer"
)

//Kafka producers to produce to different topics
//Can use one producer and produce to multiple topics but
//wanted to split work between two different producers
var stockProducer = kafkaproducer.NewDefaultProducer()
var newsProducer = kafkaproducer.NewDefaultProducer()

//StockQuoteJob defines the Job that will go and
//fetch the StockQuote for the specified company
//and output the result to kafka
type StockQuoteJob struct{
	routinepool.Job
	Company string
}

//Run executes the workflow for getting a stock quote for a given company
func (s *StockQuoteJob) Run(){
	if err := s.getStock(); err != nil{
		//return error
		s.OutputQueue <- StockQuoteOutput{Output:routinepool.Output{ID:s.ID,Success:false,Finish:time.Now()},
			Company:s.Company}
	}
	//report success
	s.OutputQueue <- StockQuoteOutput{Output:routinepool.Output{ID:s.ID,Success:true,Finish:time.Now()},
			Company:s.Company}
}

//getStock gets stock from REST endpoint and puts it in kafka queue
func (s *StockQuoteJob) getStock() error{
	stockData, err := FetchStockQuote(s.Company)
	if err != nil{
		return err
	}
	stockData.CompanyName = s.Company
	stockProducer.Input() <- &sarama.ProducerMessage{
		Topic: "stockQuotes",
		Value:stockData,
	}
	return nil
}

//StockQuoteOutput Returns the output
type StockQuoteOutput struct {
	routinepool.Output
	Company string
}


func(s StockQuoteOutput) Result() interface{}{
	return s
}
