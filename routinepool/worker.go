package routinepool

//Worker represents a worker routine that takes jobs
//from a work queue and processes the job
type Worker struct {
	WorkerChan chan Executable
	WorkQueue chan chan Executable
	QuitChan chan bool
}

//NewWorker initalizes a worker and adds a channel
//to the dispatchers workQueue that
func NewWorker(workQueue chan chan Executable) *Worker{
	workerChan := make(chan Executable)
	quitChan := make(chan bool)
	worker := Worker{WorkerChan:workerChan,
	WorkQueue:workQueue,QuitChan:quitChan}
	return &worker
}

//Start spins up worker that waits to be given job,
//processes it, then waits for another job to come in.
//Registers with the work queue by sending its channel
//into the work queue
func(w *Worker) Start(){
	go func(){
		for{
			//register with the dispatcher
			w.WorkQueue <- w.WorkerChan
			//receive a job to run
			select {
			case job := <-w.WorkerChan:
				job.Run()
			case <- w.QuitChan:
				return
			}
		}
	}()
}