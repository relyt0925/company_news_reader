package routinepool


type Worker struct {
	WorkerChan chan Executable
	WorkQueue chan chan Executable
	QuitChan chan bool
}

func NewWorker(workQueue chan chan Executable) *Worker{
	workerChan := make(chan Executable)
	quitChan := make(chan bool)
	worker := Worker{WorkerChan:workerChan,
	WorkQueue:workQueue,QuitChan:quitChan}
	return &worker
}

func(w *Worker) Start(){
	go func(){
		for{
			w.WorkQueue <- w.WorkerChan
			select {
			case job := <-w.WorkerChan:
				job.Run()
			case <- w.QuitChan:
				return
			}
		}
	}()
}