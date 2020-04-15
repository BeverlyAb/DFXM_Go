package main
// https://golang.org/ref/spec#ImportPath
//https://blog.golang.org/pipelines (multi-listener chan)
// https://flaviocopes.com/golang-event-listeners/

import (
    "github.com/klauspost/cpuid"
    "time"
    "task"
    "data"
    "depchantable"
)

func main(){
    TID_DEPTable := make(map[int][]int)
	// TID_ChanTable := make(map[int]chan data.Data)
	test := depchantable.DepChanTable{TID_DEPTable,TID_ChanTable}
	test.SET_TID_DEPTable()
	test.PrintTIDTable()
	// test.SET_TID_ChanTable()
	// test.PrintChanTable()



    t0 :=  task.Task{0, CreateDepVec(0), cpuid.CPU.LogicalCPU(),0,TID_DEPTable,TID_ChanTable}
    t1 := task.Task{1, CreateDepVec(1), cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t2 := task.Task{2, CreateDepVec(1), cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t3 := task.Task{3, CreateDepVec(2), cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t4 := task.Task{4, CreateDepVec(2), cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    // t5 := task.Task{5, CreateDepVec(2), cpuid.CPU.LogicalCPU(),2,TID_DEPTable,TID_ChanTable}
    // t6 := task.Task{6, CreateDepVec(2), cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    // t7 := task.Task{7, CreateDepVec(2), cpuid.CPU.LogicalCPU(),2,TID_DEPTable,TID_ChanTable}
    // t8 := task.Task{8, CreateDepVec(2), cpuid.CPU.LogicalCPU(),4,TID_DEPTable,TID_ChanTable}




    // c0 := make(chan data.Data)
    // c1 := make(chan data.Data)

    go t0.Compute(test.TID_ChanTable[0])
    
    go t1.Receive(test.TID_ChanTable[0])
    go t2.Receive(test.TID_ChanTable[0])
    go t3.Receive(test.TID_ChanTable[0])

    go t1.Compute(test.TID_ChanTable[1])
    go t2.Compute(test.TID_ChanTable[2])
    go t3.Compute(test.TID_ChanTable[3])

    go t4.Receive(test.TID_ChanTable[1])
    go t4.Receive(test.TID_ChanTable[2])


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
