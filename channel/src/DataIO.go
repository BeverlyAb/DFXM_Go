package main

import (
    "fmt"
    "github.com/klauspost/cpuid"
    "time"
)

type Data struct {
    msg int
    TID int
    countID  int
}
//creates data dep. of up to size, dep. are from T0 to T(size) -1
func createDepVec(size int)[]Data{
    var dataDepVec = make([]Data,size)
    for i  := 0; i < size; i++ {
        dataDepVec[i] = Data{0, i, 0}
    }
    return dataDepVec
}
func main(){
    
    t0 :=  Task{0, createDepVec(0), 0, cpuid.CPU.LogicalCPU(),0}
    t1 := Task{1, createDepVec(1), 0,cpuid.CPU.LogicalCPU(),1}
    t2 := Task{2, createDepVec(2), 0,cpuid.CPU.LogicalCPU(),2}

    // printTask(t0)
    // printTask(t1)
    // printTask(t2)

    c0 := make(chan Data)
    c1 := make(chan Data)

    go t0.compute(c0)
    go t1.receive(c0)

    go t1.compute(c1)
    go t2.receive(c1)

    time.Sleep(100*time.Millisecond)
}

func printTask(t Task){
    fmt.Println("TID ", t.TID, "CountID ",t.dataOutCount, "LogicalCPU ",cpuid.CPU.LogicalCPU())
}

//------------------------------Task Functions----------------------------------------
type Task struct{
	TID int
	dataDepVec[] Data 
    dataOutCount int//dataOutVec[] Data
    PID int
    depCount int
} 

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []Data, x int) int {
    for i, n := range a {
        if x == n.TID {
            return i
        }
    }
    return len(a)
}

//checks chan, if senderCountID >= recCountID, update data
func (t Task)receive(c chan Data){
    for recData := range c {
        msg := recData.msg
        senderID := recData.TID
        senderCountID := recData.countID

        i := Find(t.dataDepVec, senderID) 
        if i != len(t.dataDepVec){
            if senderCountID == t.dataDepVec[i].countID{
                t.dataDepVec[i].msg = msg
                t.depCount++
                fmt.Println("TID ", t.TID, "CountID ",t.dataOutCount, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.dataDepVec[i].msg)
                 fmt.Println("TID ", t.TID, "CountID ",t.dataOutCount, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.dataDepVec[i].msg)
                  fmt.Println("TID ", t.TID, "CountID ",t.dataOutCount, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.dataDepVec[i].msg)
 fmt.Println("TID ", t.TID, "CountID ",t.dataOutCount, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.dataDepVec[i].msg)

            }
        }
    }
}

//checks if all dep. met
func (t Task)readyToCompute()bool{
    return len(t.dataDepVec) == t.depCount
}

//computes and fire
func (t Task)compute(c chan Data){
    if t.readyToCompute() {
        t.depCount = 0

        var msg int = (t.TID + 1) *100
        var TID int = t.TID
        var countID int = t.dataOutCount
        var dataOut Data = Data{msg, TID, countID}
        t.fire(dataOut, c)
    }
}

//opens chan 
func (t Task)fire(dataOut Data, c chan Data){
    defer close(c)
    c<-dataOut
}
