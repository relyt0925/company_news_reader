package routinepool

import (
	"time"
	"github.com/relyt0925/company_news_reader/myallies"
	"github.com/relyt0925/company_news_reader/kafkaproducer"
	"github.com/Shopify/sarama"

	//"sync/atomic"
)

var stockProducer = kafkaproducer.NewDefaultProducer()
var newsProducer = kafkaproducer.NewDefaultProducer()


type Submitter struct {
	workQueue chan<- Executable
	outputQueue chan ExecutableResult
	outputKill chan bool
	jobCounter int32
	outputHandler OutputHandler
}

type ExecutableResult interface {
	Result() interface{}
}

type OutputHandler func(<-chan ExecutableResult, <-chan  bool)

func NewSubmitter(workQueueSize int, workerPoolSize int, outputQueueSize int, fn OutputHandler) *Submitter{
	workQueue := make(chan Executable,workQueueSize)
	outputQueue := make(chan ExecutableResult, outputQueueSize)
	outputKill := make(chan bool)
	startDispatcher(workerPoolSize,workQueue)
	submitter := Submitter{workQueue:workQueue,outputQueue:outputQueue,jobCounter:0,outputHandler:fn,outputKill:outputKill}
	go submitter.outputHandler(submitter.outputQueue,submitter.outputKill)
	return &submitter
}


func(s *Submitter) Submit(job Executable){
	//job.SetID(atomic.AddInt32(&s.jobCounter,1))
	job.SetOutputChan(s.outputQueue)
	s.workQueue <- job
}

type Executable interface{
	SetID(int32)
	Run()
	SetOutputChan(chan<- ExecutableResult)
}


type Job struct {
	ID int32
	Args []interface{}
	OutputQueue chan<- ExecutableResult
}

func (j *Job) SetID (ID int32){
	j.ID=ID
}

func (j *Job) SetOutputChan(channel chan<- ExecutableResult){
	j.OutputQueue=channel
}


type StockQuoteJob struct{
	*Job
	Company string

}

func (s StockQuoteJob) Run(){
	if err := s.getStock(); err != nil{
		//return error
		s.OutputQueue <- StockQuoteOutput{Output:&Output{ID:s.ID,Success:false,Finish:time.Now()},
			Company:s.Company}
	}
	//report success
	s.OutputQueue <- StockQuoteOutput{Output:&Output{ID:s.ID,Success:true,Finish:time.Now()},
			Company:s.Company}
}

func (s StockQuoteJob) SetID(ID int32){
	s.ID=ID
}

func (s StockQuoteJob) getStock() error{
	stockData, err := myallies.FetchStockQuote(s.Company)
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

type Output struct{
	ID int32
	Success bool
	Finish time.Time
}

type StockQuoteOutput struct {
	*Output
	Company string
}


func(s StockQuoteOutput) Result() interface{}{
	return s
}








