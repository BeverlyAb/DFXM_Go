package main
// https://golang.org/ref/spec#ImportPath
import (
    "github.com/klauspost/cpuid"
    "time"
    "task"
    "data"
    "depchantable"
)

func main(){
    TID_DEPTable := make(map[int][]int)
	TID_ChanTable := make(map[int]chan data.Data)
	test := depchantable.DepChanTable{TID_DEPTable,TID_ChanTable}
	test.SET_TID_DEPTable()
	test.PrintTIDTable()
	test.SET_TID_ChanTable()
	test.PrintChanTable()



    t0 :=  task.Task{0, test.TID_DEPTable[0], cpuid.CPU.LogicalCPU(),0,TID_DEPTable,TID_ChanTable}
    t1 := task.Task{1, test.TID_DEPTable[1], cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t2 := task.Task{2, test.TID_DEPTable[2], cpuid.CPU.LogicalCPU(),2,TID_DEPTable,TID_ChanTable}


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
