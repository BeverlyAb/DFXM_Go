package task

import (
    "fmt"
    "github.com/klauspost/cpuid"
    "data"
)

type Task struct{
	TID int
	DataDepVec[] data.Data 
    //DataOutMsg int//dataOutVec[] Data
    PID int
    DepCount int
    TID_DEPTable map[int][] int
    TID_ChanTable map[int]chan data.Data
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
    if t.DiscernChan(c){
        for recData := range c {
            msg := recData.Msg
            senderID := recData.TID
            senderCountID := recData.CountID

            i := Find(t.DataDepVec, senderID) 
            if i != len(t.DataDepVec){
                if senderCountID == t.DataDepVec[i].CountID{
                    t.DataDepVec[i].Msg = msg
                    t.DepCount++
                    fmt.Println("Recvd:\tTID ", t.TID,  "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.DataDepVec[i].Msg,"Msg ID ", senderCountID)
                    // if cpuid.CPU.LogicalCPU() % 2 ==0{
                    //     fmt.Println("even ",cpuid.CPU.LogicalCPU())
                    // }
                }
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
    if t.ReadyToCompute(){
        t.DepCount = 0

        var msg int = (t.TID + 1) *100
        var TID int = t.TID
        //var countID int = t.DataOutMsg
        var countID = 0
        var dataOut data.Data = data.Data{msg, TID, countID}
        t.Fire(dataOut, c)
        fmt.Println("Fired:\tTID ", TID, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", dataOut.Msg, "Msg ID ", dataOut.CountID)
    }
}

//opens chan 
func (t Task)Fire(dataOut data.Data, c chan data.Data){
    defer close(c)
    c<-dataOut
}

//returns true if task should listen on the channel, since it contains its dep
func(t Task)DiscernChan(chanIn chan data.Data)bool{
    var vecTID = make([]int,4)
    vecTID = t.TID_DEPTable[t.TID]

    for TID := range vecTID{
        if t.TID_ChanTable[TID] == chanIn{
            fmt.Println("TID ",t.TID, " SenderTID ", TID)
            return true
        }
    }
    return false
}


func PrintTask(t Task){
    //fmt.Println("TID ", t.TID, "CountID ",t.DataOutMsg, "LogicalCPU ",cpuid.CPU.LogicalCPU())
}