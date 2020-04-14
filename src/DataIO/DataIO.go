package main
// https://golang.org/ref/spec#ImportPath
import (
    "github.com/klauspost/cpuid"
    "time"
    "task"
    "data"
)

func main(){
    
    t0 :=  task.Task{0, CreateDepVec(0), cpuid.CPU.LogicalCPU(),0}
    t1 := task.Task{1, CreateDepVec(1), cpuid.CPU.LogicalCPU(),1}
    t2 := task.Task{2, CreateDepVec(2), cpuid.CPU.LogicalCPU(),2}

    // printTask(t0)
    // printTask(t1)
    // printTask(t2)

    c0 := make(chan data.Data)
    c1 := make(chan data.Data)

    go t0.Compute(c0)
    go t1.Receive(c0)

    go t1.Compute(c1)
    go t2.Receive(c1)

    time.Sleep(100*time.Millisecond)
}


//creates data dep. of up to size, dep. are from T0 to T(size) -1
func CreateDepVec(size int)[]data.Data{
    var dataDepVec = make([]data.Data,size)
    for i  := 0; i < size; i++ {
        dataDepVec[i] = data.Data{0, i, 0}
    }
    return dataDepVec
}
