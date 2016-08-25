package routinepool

import (
	"time"
	"sync/atomic"
)


//Submitter provides framework to submit jobs to pool of workers
//and have workers run jobs then submit the jobs output to be
//processed by a handler function
type Submitter struct {
	workQueue chan<- Executable
	outputQueue chan ExecutableResult
	outputKill chan bool
	jobCounter int32
	outputHandler OutputHandler
}

//ExecuableResult provides interface to be able
//to fetch result of a Executable
type ExecutableResult interface {
	Result() interface{}
}

//OutputHandler defines message signature required by a output handler
type OutputHandler func(<-chan ExecutableResult, <-chan  bool)

//NewSubmitter creates a New Submitter with the specified work queue size, output queue size, worker pool size,
// and output handler function
func NewSubmitter(workQueueSize int, workerPoolSize int, outputQueueSize int, fn OutputHandler) *Submitter{
	workQueue := make(chan Executable,workQueueSize)
	outputQueue := make(chan ExecutableResult, outputQueueSize)
	outputKill := make(chan bool)
	startDispatcher(workerPoolSize,workQueue)
	submitter := Submitter{workQueue:workQueue,outputQueue:outputQueue,jobCounter:0,outputHandler:fn,outputKill:outputKill}
	go submitter.outputHandler(submitter.outputQueue,submitter.outputKill)
	return &submitter
}

//Submit submits the job to the worker pool
func(s *Submitter) Submit(job Executable){
	job.SetID(atomic.AddInt32(&s.jobCounter,1))
	job.SetOutputChan(s.outputQueue)
	s.workQueue <- job
}

//Executable defines interface to specify the job that is being run
type Executable interface{
	SetID(int32)
	Run()
	SetOutputChan(chan<- ExecutableResult)
}

//Job provides base struct to embed in other structs and
//extend to create Executable interface
type Job struct {
	ID int32
	OutputQueue chan<- ExecutableResult
}

//SetID sets the job id for the job
func (j *Job) SetID (ID int32){
	j.ID=ID
}

//SetOutputChan sets output channel that job results will
//be sent on
func (j *Job) SetOutputChan(channel chan<- ExecutableResult){
	j.OutputQueue=channel
}

//Output provides base struct to embed in other structs
// and create result interface
type Output struct{
	ID int32
	Success bool
	Finish time.Time
}

//Result just returns the output of the job
func(o Output) Result() interface{}{
	return o
}








