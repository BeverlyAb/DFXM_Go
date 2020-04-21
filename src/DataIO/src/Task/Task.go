package task

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
    return data.Data{msg, t.TID, countID}
}

// //creates a channel for every type of data
// func (t * Task)Producer(iters int) <-chan data.Data {
//     c := make(chan data.Data)
//     go func() {
//         for i := 0; i < iters; i++ {
//             msg := int(math.Pow(10,float64(t.TID)))
//             c <- data.Data{msg,t.TID,0}
//         //    fmt.Println("TID ", t.TID, "Fired and Converted ",msg) 
//         }
//         //t.Invalidate()  
//         close(c)
//     }()
    
//     return c
// }

// //fans out channels which stem from firing tasks
// //# channels = size of Send_to 
// func (t * Task)FanOutUnbuffered(ch <-chan data.Data, size int) []chan data.Data {
//     cs := make([]chan data.Data, size)
//     for i, _ := range cs {
//         // The size of the channels buffer controls how far behind the recievers
//         // of the fanOut channels can lag the other channels.
//         cs[i] = make(chan data.Data)
//     }
//     go func() {
//         for i := range ch {
//             for _, c := range cs {
//                 c <- i
//             }
//         }
//         for _, c := range cs {
//             // close all our fanOut channels when the input channel is exhausted.
//             close(c)
//         }
//     }()
//     return cs
// }


//assigns the channels on the receiving end after a task fires
//and receivers consume data
func (t * Task)assignRecChan(chanSet []chan data.Data, TaskSet *  [] Task){
     
        for i := 0; i < len(chanSet); i++ {
           go (*TaskSet)[t.Send_to[i]].Consumer(chanSet[i], t.TID)
        }

}

//consumes data from channel and updates RecFrom
func (t * Task)Consumer(cin <-chan data.Data, senderTID int) {
    //go func(){
        for recData := range cin {
             t.updateRecFrom(senderTID)
            // fmt.Println("TID ",t.TID, "Sent data to ",t.Send_to[i])
           fmt.Println("TID ", t.TID, " Received  ",recData.Msg, " from TID ",recData.TID)
        }
    //}()
}

//updates RecFrom from recevier view to indicate that the sender sent data and was received
func (t * Task)updateRecFrom(senderTID int){
    t.Rec_from[senderTID] = true
}

// //changes TID to invalidate
// func (t * Task)Invalidate(){
//     t.TID = -1
// }

