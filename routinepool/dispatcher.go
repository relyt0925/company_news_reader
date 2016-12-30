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
			case work, ok := <-workQueue:
				// Closing the channel kills the dispatcher and cleans up goroutines
				if !ok{
					for i:=0 ; i<numWorkers; i++ {
						w := <- workerPool
						close(w)
					}
					close(workerPool)
					return
				}
				worker := <- workerPool
				go func() {
					worker <- work
				}()
			}
		}
	}()
}
