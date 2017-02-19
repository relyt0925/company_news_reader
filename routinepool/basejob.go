package routinepool


//Job provides base struct to embed in other structs and
//extend to create Executable interface
type Job struct {
	OutputQueue chan<- ExecutableResult
}


//ExecuableResult provides interface to be able
//to fetch result of a Executable
type ExecutableResult interface {
	Result() (interface{}, error)
}

//Output provides base struct to embed in other structs
// and create result interface
type Output struct{
	Res interface{}
	Err error
}

//Result just returns the output of the job
func(o Output) Result() (interface{}, error){
	if o.Err != nil{
		return nil, o.Err
	}
	return o.Res, nil
}
