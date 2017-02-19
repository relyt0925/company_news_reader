package routinepool

import (
	"testing"
	"fmt"
	"time"
)

func TestNewWorker(t *testing.T) {

	testWorkQueue := make(chan chan Executable, 5)
	w := NewWorker(testWorkQueue)
	if w.WorkerChan == nil {
		t.Fatal("Worker Channel is nil.")
	}
	if w.WorkQueue != testWorkQueue{
		t.Fatal("Workqueue does not eqaul test work queue.")
	}
}

func TestWorker_Start(t *testing.T) {
	testWorkers := 5
	testWorkQueue := make(chan chan Executable, testWorkers)
	for i:=0; i<testWorkers; i++{
		w := NewWorker(testWorkQueue)
		w.Start()

	}
	//give them time to join channel
	time.Sleep(1*time.Millisecond)
	if len(testWorkQueue) != testWorkers{
		t.Fatal(" All Workers Did Not Spin Up Properly. Length",len(testWorkQueue))
	}
	for i := 0; i<10; i++{
		i := testWorkerJob{}
		w := <- testWorkQueue
		w <- i
	}
	//give time for jobs to process and rejoin channel
	time.Sleep(1*time.Millisecond)
	if len(testWorkQueue) != testWorkers{
		t.Fatal(" All Workers Did Not Rejoin work queue after running job. Length",len(testWorkQueue) )
	}
}


type testWorkerJob struct {
}


func (t testWorkerJob) Run(){
	fmt.Println("Running!")
}
