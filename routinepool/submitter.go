package routinepool

import "sync"

//Submitter provides framework to submit jobs to pool of workers
//and have workers run jobs then submit the jobs output to be
//processed by a handler function
type Submitter struct {
	workQueue chan<- Executable
	closed bool
	closedMutex sync.Mutex
}

//Executable defines interface to specify the job that is being run
type Executable interface{
	Run()
}


//NewSubmitter creates a New Submitter with the specified work queue size, output queue size, worker pool size,
// and output handler function
func NewSubmitter(workQueueSize int, workerPoolSize int) *Submitter{
	workQueue := make(chan Executable,workQueueSize)
	startDispatcher(workerPoolSize,workQueue)
	submitter := Submitter{workQueue:workQueue,closed:false,closedMutex:sync.Mutex{}}
	return &submitter
}

//Submit submits the job to the worker pool
func(s *Submitter) Submit(job Executable){
	s.workQueue <- job
}

//Cleanup shuts down all worker threads after they complete their work
func(s *Submitter) Cleanup() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()
	if !s.closed{
		close(s.workQueue)
		s.closed=true
	}
}










