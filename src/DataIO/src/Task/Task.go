package task

import (
    "fmt"
    "github.com/klauspost/cpuid"
    "data"
)

// type Data struct {
//     msg int
//     TID int
//     countID  int
// }

type Task struct{
	TID int
	DataDepVec[] data.Data 
    DataOutCount int//dataOutVec[] Data
    PID int
    DepCount int
} 

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []data.Data, x int) int {
    for i, n := range a {
        if x == n.TID {
            return i
        }
    }
    return len(a)
}

//checks chan, if senderCountID >= recCountID, update data
func (t Task)Receive(c chan data.Data){
    for recData := range c {
        msg := recData.Msg
        senderID := recData.TID
        senderCountID := recData.CountID

        i := Find(t.DataDepVec, senderID) 
        if i != len(t.DataDepVec){
            if senderCountID == t.DataDepVec[i].CountID{
                t.DataDepVec[i].Msg = msg
                t.DepCount++
               fmt.Println("TID ", t.TID, "CountID ",t.DataOutCount, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.DataDepVec[i].Msg)
            }
        }
    }
}

//checks if all dep. met
func (t Task)ReadyToCompute()bool{
    return len(t.DataDepVec) == t.DepCount
}

//computes and fire
func (t Task)Compute(c chan data.Data){
    if t.ReadyToCompute() {
        t.DepCount = 0

        var msg int = (t.TID + 1) *100
        var TID int = t.TID
        var countID int = t.DataOutCount
        var dataOut data.Data = data.Data{msg, TID, countID}
        t.Fire(dataOut, c)
    }
}

//opens chan 
func (t Task)Fire(dataOut data.Data, c chan data.Data){
    defer close(c)
    c<-dataOut
}

func PrintTask(t Task){
    fmt.Println("TID ", t.TID, "CountID ",t.DataOutCount, "LogicalCPU ",cpuid.CPU.LogicalCPU())
}