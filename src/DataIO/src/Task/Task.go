package task
//https://medium.com/a-journey-with-go/go-concurrency-access-with-maps-part-iii-8c0a0e4eb27e
//concurrent map writes when timeTrack is removed
import (
    "fmt"
    // "github.com/klauspost/cpuid"
    "data"
    "math"
    // "time"
    //"depchantable"
    "fans"
)

type Task struct{
    TID int
    Rec_from map[int]bool   //TID_rec : Received?
    Send_to []int           //list of TID to send to
    PID int                 //Processor ID
    DataRecvd [] data.Data  //list of Data received 
   // DataChan chan data.Data //communication channel
} 

//Tasks sends data and releases info
func (t * Task)Fire(iters int, TaskSet * [] Task) {
    if t.readyToCompute(){
        done := make(chan bool)
        defer close(done)
        var buffer int = 1
        var fanOutSize int = len(t.Send_to)
        chanSet := fans.FanOut(done,buffer,fanOutSize,t.ComputeAndProduce())


        t.assignRecChan(chanSet, TaskSet)
    }   
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

//compute some Task and produce single output
func (t * Task)ComputeAndProduce()data.Data{
    var msg int
    for i := 0; i < len(t.DataRecvd); i++ {
        msg = int(math.Pow(10,float64(t.TID)))*t.DataRecvd[i].Msg
    }
    countID := 0
    fmt.Println("TID ",t.TID," Fired")
    return data.Data{msg, t.TID, countID}
}

//assigns the channels on the receiving end after a task fires
//and receivers consume data
func (t * Task)assignRecChan(chanSet []chan data.Data, TaskSet *  [] Task){
     
        for i := 0; i < len(chanSet); i++ {
           go (*TaskSet)[t.Send_to[i]].updateRecFrom(t.TID)
        }

}

//consumes data from channel and updates RecFrom
// func (t * Task)Consumer(cin <-chan data.Data, senderTID int) {
//     t.updateRecFrom(senderTID)
// }

//updates RecFrom from recevier view to indicate that the sender sent data and was received
func (t * Task)updateRecFrom(senderTID int){
    t.Rec_from[senderTID] = true
}

// //changes TID to invalidate
// func (t * Task)Invalidate(){
//     t.TID = -1
// }

