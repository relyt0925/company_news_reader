package routinepool

import (
	"github.com/bouk/monkey"
	"testing"
	"fmt"
	"time"
)

//Be sure to run this test with a timeout value because it tests that the mutex
//is initalized properly
func TestNewSubmitter(t *testing.T){
	var testSubmitterWorkQueueSize int = 5
	var testSubmitterWorkerPoolSize int = 5
	var testSubmitterfailure bool = false
	monkey.Patch(startDispatcher, func(numWorkers int, workQueue chan Executable){
		if numWorkers != testSubmitterWorkerPoolSize || cap(workQueue) != testSubmitterWorkQueueSize{
			fmt.Println("numworkers",numWorkers)
			fmt.Println("testSubmitterWorkerPoolSize",testSubmitterWorkerPoolSize)
			fmt.Println("workQueue",cap(workQueue))
			fmt.Println("testSubmitterWorkQueueSize",testSubmitterWorkQueueSize)
			testSubmitterfailure = true
		}
	})
	defer monkey.UnpatchAll()
	submitter := NewSubmitter(testSubmitterWorkerPoolSize,testSubmitterWorkQueueSize)
	if testSubmitterfailure{
		t.Fatal("Worker Pool Size or Worker Queue Size Did Not Match Expected Value")
	}
	if submitter.closed{
		t.Fatal("Closed should be initalized to false")
	}
	submitter.closedMutex.Lock()
	submitter.closedMutex.Unlock()
}

func TestSubmitter_Submit(t *testing.T) {
	testWorkQueue := make(chan Executable, 5)
	testSubmitter := Submitter{workQueue:testWorkQueue}
	testJob := testSubmitterJob{i:10}
	testSubmitter.Submit(testJob)
	receivedJob := <- testWorkQueue
	s, ok := receivedJob.(testSubmitterJob)
	if !ok || s.i != testJob.i{
		t.Fatal("Received Job is not same as Submitted Job")
	}
}

func TestSubmitter_SubmitConcurrent(t *testing.T) {
	testWorkQueue := make(chan Executable, 5)
	testSubmitter := Submitter{workQueue:testWorkQueue}
	testJob := testSubmitterJob{i:10}
	numJobs := 10
	for i:= 0; i<numJobs; i++{
		go func(t *Submitter){
			t.Submit(testJob)
		}(&testSubmitter)
	}
	for i:= 0; i<numJobs; i++{
		receivedJob := <- testWorkQueue
		s, ok := receivedJob.(testSubmitterJob)
		if !ok || s.i != testJob.i{
			t.Fatal("Received Job is not same as Submitted Job")
		}
	}
}

func TestSubmitter_Cleanup(t *testing.T) {
	testWorkQueue := make(chan Executable, 5)
	testSubmitter := Submitter{workQueue:testWorkQueue}
	for i:=0; i<2; i++{
		go func(submitter *Submitter){
			testSubmitter.Cleanup()
		}(&testSubmitter)
	}
	//to ensure both routines try and close the channel
	time.Sleep(1*time.Second)
	_, stillOpen := <- testWorkQueue
	if stillOpen{
		t.Fatal("Received a job when channel should be closed")
	}
}

type testSubmitterJob struct {
	i int
}


func (t testSubmitterJob) Run(){
	fmt.Println("Running!")
}