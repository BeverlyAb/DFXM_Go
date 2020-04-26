package task
//https://medium.com/a-journey-with-go/go-concurrency-access-with-maps-part-iii-8c0a0e4eb27e
//concurrent map writes when timeTrack is removed
import (
    // "fmt"
    // "github.com/klauspost/cpuid"
    "data"
    "math"
    "time"
    "fans"
    "timeout"
    "faulthandler"
)

type Task struct{
    TID int
    Rec_from map[int]bool   //TID_rec : Received?
    Send_to []int           //list of TID to send to
    PID int                 //Processor ID
    DataRecvd [] data.Data  //list of Data received
    Timeout time.Duration 
    RecomputeCount int         //keeps track of recompute, used for Data's CountID
} 

//Tasks sends data and releases info; returns bool to keep track of tasks
//that still needs to fire
func (t * Task)Fire(TaskSet * [] Task, FH * faulthandler.FaultHandler) bool{
    if t.readyToCompute(){
        done := make(chan bool)
        defer close(done)

        var buffer int = 1
        var fanOutSize int = len(t.Send_to)
        dataOut, recompute := t.ComputeAndProduce(FH)
        chanSet:= fans.FanOut( done,buffer,
                                        fanOutSize,
                                        recompute, 
                                        dataOut)
        if recompute {
            t.Timeout *= 2
            return false
        } else {
            t.assignRecChan(chanSet, TaskSet)
            return true
        }
    }   
    return false
}

//checks if task is ready to fire if all dependencies met
func (t * Task)readyToCompute()bool{
    for _,element := range t.Rec_from{
        if !element {
            return false
        }
    }
    return true
}

//compute some Task and produce single output
func (t * Task)ComputeAndProduce(FH * faulthandler.FaultHandler)(data.Data,bool){
    tm := new(timeout.Timeout)
    tm.Init(t.Timeout)

    //inserting long computation
    FH.InsertRecompute(t.TID)

    var msg int
    //include !dead = dead will always be true
    for i := 0; i < len(t.DataRecvd) && !tm.HasTimedOut; i++ {
        tm.Update() 
        //do computation based on received data individually
        msg += t.DataRecvd[i].Msg
    }

    msg += int(math.Pow(10,float64(t.TID)))
    if tm.HasTimedOut{
        t.RecomputeCount++
    }
    countID := t.RecomputeCount
    return data.Data{msg, t.TID, countID}, tm.HasTimedOut
}

//assigns the channels on the receiving end after a task fires
//and receivers consume data
func (t * Task)assignRecChan(chanSet []chan data.Data, TaskSet *  [] Task){
     
        for i := 0; i < len(chanSet); i++ {
            t.recData(&(*TaskSet)[t.Send_to[i]], chanSet[i])
            (*TaskSet)[t.Send_to[i]].updateRecFrom(t.TID)
        }

}

//updates RecFrom from recevier view to indicate that the sender sent data and was received
func (t * Task)updateRecFrom(senderTID int){
    t.Rec_from[senderTID] = true
}

//add data from sender to recevier's RecData
func (t * Task)recData(receiver * Task, in chan data.Data){
    for elem := range in {
        receiver.DataRecvd = append(receiver.DataRecvd, elem) 
    } 
}