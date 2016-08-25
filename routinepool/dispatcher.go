package routinepool

//startDispatcher spins off worker threads and handles passing data
//from the work queue to the worker threads.
func startDispatcher(numWorkers int, workQueue chan Executable){
	workerPool := make(chan chan Executable, numWorkers)
	for i := 0; i<numWorkers; i++ {
		worker := NewWorker(workerPool)
		worker.Start()
	}
	go func(){
		for{
			select {
			case work := <-workQueue:
				worker := <- workerPool
				go func() {
					worker <- work
				}()
			}
		}
	}()
}
