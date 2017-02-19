package routinepool

import "fmt"

//Worker represents a worker routine that takes jobs
//from a work queue and processes the job
type worker struct {
	WorkerChan chan Executable
	WorkQueue chan chan Executable
}

//NewWorker initalizes a worker and adds a channel
//to the dispatchers workQueue that
func NewWorker(workQueue chan chan Executable) *worker{
	workerChan := make(chan Executable)
	worker := worker{WorkerChan:workerChan,
	WorkQueue:workQueue}
	return &worker
}

//Start spins up worker that waits to be given job,
//processes it, then waits for another job to come in.
//Registers with the work queue by sending its channel
//into the work queue
func(w *worker) Start(){
	go func(){
		for{
			//register with the dispatcher
			w.WorkQueue <- w.WorkerChan

			//receive a job to run
			select {
			case job, ok := <-w.WorkerChan:
				if !ok {
					fmt.Println("CLOSING")
					return
				}
				job.Run()
			}
		}
	}()
}