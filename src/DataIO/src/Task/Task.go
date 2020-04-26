package task
//https://medium.com/a-journey-with-go/go-concurrency-access-with-maps-part-iii-8c0a0e4eb27e
//concurrent map writes when timeTrack is removed
import (
    // "fmt"
    // "github.com/klauspost/cpuid"
    "data"
    "math"
    "time"
    //"depchantable"
    "fans"
)

type Task struct{
    TID int
    Rec_from map[int]bool   //TID_rec : Received?
    Send_to []int           //list of TID to send to
    PID int                 //Processor ID
    DataRecvd [] data.Data  //list of Data received
    Timeout time.Duration 
    RefireCount int         //keeps track of refire, used for Data's CountID
} 

//Tasks sends data and releases info; returns bool to keep track of tasks
//that still needs to fire
func (t * Task)Fire(TaskSet * [] Task) bool{
    if t.readyToCompute(){
        done := make(chan bool)
        defer close(done)

        var buffer int = 1
        var fanOutSize int = len(t.Send_to)
        dataOut, refire := t.ComputeAndProduce()
        chanSet:= fans.FanOut( done,buffer,
                                        fanOutSize,
                                        refire, 
                                        dataOut)
        if refire {
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
func (t * Task)ComputeAndProduce()(data.Data,bool){
    var msg int
    start := time.Now()
    elapsed := time.Now().Sub(start)

   //inserting long computation or busy computation
   // if t.TID == 5 {
   //      time.Sleep(time.Millisecond*70)
   //  }
    for i := 0; i < len(t.DataRecvd) && elapsed < t.Timeout; i++ {
        elapsed = time.Now().Sub(start) 
    
        //do computation based on received data individually
        //if t.DataRecvd[i].CountID == t.CountID{ //create matching countID
            msg += t.DataRecvd[i].Msg
        //}
    }
    msg += int(math.Pow(10,float64(t.TID)))
    slack := time.Millisecond*10
    hasTimedOut :=  elapsed > t.Timeout + slack
    
    if hasTimedOut{
        t.RefireCount++
        //notify receivers of my countID
    }
    countID := t.RefireCount
    
    return data.Data{msg, t.TID, countID}, hasTimedOut
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

func (t * Task)recData(receiver * Task, in chan data.Data){
    for elem := range in {
        receiver.DataRecvd = append(receiver.DataRecvd, elem) 
    } 
}