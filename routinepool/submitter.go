package routinepool


//Submitter provides framework to submit jobs to pool of workers
//and have workers run jobs then submit the jobs output to be
//processed by a handler function
type Submitter struct {
	workQueue chan<- Executable
}



//OutputHandler defines message signature required by a output handler
type OutputHandler func(<-chan ExecutableResult, <-chan  bool)


//NewSubmitter creates a New Submitter with the specified work queue size, output queue size, worker pool size,
// and output handler function
func NewSubmitter(workQueueSize int, workerPoolSize int) *Submitter{
	workQueue := make(chan Executable,workQueueSize)
	startDispatcher(workerPoolSize,workQueue)
	submitter := Submitter{workQueue:workQueue}
	return &submitter
}

//Submit submits the job to the worker pool
func(s *Submitter) Submit(job Executable){
	s.workQueue <- job
}

func(s *Submitter) Cleanup() {
	close(s.workQueue)
}


//Executable defines interface to specify the job that is being run
type Executable interface{
	Run()
}

//Job provides base struct to embed in other structs and
//extend to create Executable interface
type Job struct {
	OutputQueue chan<- ExecutableResult
}


//ExecuableResult provides interface to be able
//to fetch result of a Executable
type ExecutableResult interface {
	Result() (interface{}, error)
}

//Output provides base struct to embed in other structs
// and create result interface
type Output struct{
	Res interface{}
	Err error
}

//Result just returns the output of the job
func(o Output) Result() (interface{}, error){
	if o.Err != nil{
		return nil, o.Err
	}
	return o.Res, nil
}








