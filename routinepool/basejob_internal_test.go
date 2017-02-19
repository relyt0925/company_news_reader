package routinepool

import (
	"testing"
	"fmt"
)

func TestOutput_ResultError(t *testing.T) {
	e := fmt.Errorf("ERROR")
	o := Output{Err:e,Res:nil}
	res, err := o.Result()
	if res != nil{
		t.Fatal("Result is not nil!")
	}
	if err == nil{
		t.Fatal("Error is nil!")
	}
	if err != e{
		t.Fatal("Error returned is not expected error")
	}
}

func TestOutput_ResultValid(t *testing.T) {
	o := Output{Err:nil,Res:int(5)}
	res, err := o.Result()
	if res == nil{
		t.Fatal("Result is nil!")
	}
	if err != nil{
		t.Fatal("Error is not nil!")
	}
	i, ok := res.(int);
	if !ok{
		t.Fatal("Result is not expected type!")
	}
	if i != 5{
		t.Fatal("Result Does not have expected value!")
	}
}