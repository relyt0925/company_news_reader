package routinepool


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
