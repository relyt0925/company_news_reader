package routinepool

import (
	"testing"
	"sync/atomic"
	"fmt"
	"time"
	"github.com/bouk/monkey"
	"reflect"
)

//global so it can be seen
var closeCounter uint32

func TestStartDispatcher(t *testing.T){
	var counter uint32 = 0
	testWorkQueueSize := 20
	testWorkQueue := make(chan Executable,testWorkQueueSize)
	testNumWorkers := 10
	var w *worker
	monkey.PatchInstanceMethod(reflect.TypeOf(w),"Start", mockStart)
	defer monkey.UnpatchAll()
	startDispatcher(testNumWorkers,testWorkQueue)
	var testNumJobs uint32 = 20
	for i:=uint32(0);i<testNumJobs;i++{
		i := testDispatcherJob{counterRef:&counter}
		testWorkQueue <- i
	}
	close(testWorkQueue)
	//give jobs time to complete (remember this is asyncronous)
	time.Sleep(1*time.Second)
	if atomic.LoadUint32(&counter) != testNumJobs{
		t.Fatalf("Not All Jobs Ran Successfully. Error submitting jobs to Pool")
	}
	if atomic.LoadUint32(&closeCounter) != uint32(testNumWorkers){
		t.Fatalf("Not All Threads Stopped. Resources being leaked")
	}
}

type testDispatcherJob struct {
	counterRef *uint32
}


func (t testDispatcherJob) Run(){
	fmt.Println("Running!")
	atomic.AddUint32(t.counterRef,1)
}


//Mocks The start method to also increment closed counter
func mockStart(w *worker){
	go func(){
		for{
			//register with the dispatcher
			w.WorkQueue <- w.WorkerChan

			//receive a job to run
			select {
			case job, ok := <-w.WorkerChan:
				if !ok {
					fmt.Println("CLOSING")
					atomic.AddUint32(&closeCounter,1)
					return
				}
				job.Run()
			}
		}
	}()
}