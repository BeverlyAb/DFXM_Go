package task

import (
    "fmt"
    // "github.com/klauspost/cpuid"
    "data"
    "math"
    "time"
    //"depchantable"
)

type Task struct{
    TID int
    Rec_from map[int]bool   //TID_rec : Received?
    Send_to []int           //list of TID to send to
    PID int                 //Processor ID
   // DataChan chan data.Data //communication channel
} 

//checks if task is ready to fire if all dependencies met
func (t * Task)readyToCompute()bool{
    if len(t.Rec_from) == 0 {
        return true
    } 
    for i := 0; i < len(t.Rec_from); i++ {
        if !t.Rec_from[i]{
            return false
        } 
    }
    return true
}

//creates a channel for every type of data
func (t * Task)Producer(iters int) <-chan data.Data {
    c := make(chan data.Data)
    go func() {
        for i := 0; i < iters; i++ {
            msg := int(math.Pow(10,float64(t.TID)))
            c <- data.Data{msg,t.TID,0}
            fmt.Println("TID ", t.TID, "Fired and Converted ",msg)
            time.Sleep(1 * time.Second)
        }
        close(c)
    }()
    return c
}

func (t * Task)Consumer(cin <-chan data.Data) {
    for recData := range cin {
            
            msg := recData.Msg
            senderID := recData.TID
            senderCountID := recData.CountID


            i := Find(t.DataDepVec, senderID) 
            if i != len(t.DataDepVec){
                if senderCountID == t.DataDepVec[i].CountID{
                    t.DataDepVec[i].Msg = msg
                    t.DepCount++
                    t.DataOutMsg += msg
                    // fmt.Println("Recvd:\tTID ", t.TID,  "LogicalCPU ", cpuid.CPU.LogicalCPU(), "NewData ", t.DataOutMsg,"Msg ID ", senderCountID)
                    fmt.Println("TID ", t.TID, " Received  ",t.DataDepVec[i].Msg, " from TID ",senderID)
                }
            }


        // fmt.Println("Msg ",recData.Msg, "SenderTID ",recData.TID, "MsgCount ",recData.CountID)
    }


}

//fans out channels which stem from firing tasks
//# channels = size of Send_to 
func (t Task)FanOutUnbuffered(ch <-chan data.Data, size int) []chan data.Data {
    cs := make([]chan data.Data, size)
    for i, _ := range cs {
        // The size of the channels buffer controls how far behind the recievers
        // of the fanOut channels can lag the other channels.
        cs[i] = make(chan data.Data)
    }
    go func() {
        for i := range ch {
            for _, c := range cs {
                c <- i
            }
        }
        for _, c := range cs {
            // close all our fanOut channels when the input channel is exhausted.
            close(c)
        }
    }()
    return cs
}

//assigns the channels on the receiving end after a task fires
//and receivers consume data
func (t * Task)assignRecChan(chanSet []chan data.Data, TaskSet * [] task){
    for i := 0; i < len(t.Send_to); i++ {
        //TaskSet[t.Send_to[i]] = chanSet[i]
        TaskSet[t.Send_to[i]].Consumer(chanSet[i])
        t.updateRecFrom(t.Send_to[i])
    }
}

//Tasks sends data and releases info
func (t * Task)Fire(iters int, TaskSet * [] task) {
    if t.ReadyToCompute(){
        mainChan := t.Producer(iters)
        chanSet := t.FanOutUnbuffered(mainChan,len(t.Send_to))
        t.assignRecChan(chanSet, & TaskSet)
    }   
}


func (t * Task)updateRecFrom(index int){
    t.Rec_from[index] = true
}
// func (t Task)Init(TID int, recFrom map[int]bool, sendTo []int){
//     t.TID = TID
//     t.Rec_from = recFrom
//     t.Send_to = sendTo
// }

