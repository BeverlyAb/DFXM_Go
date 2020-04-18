package task

import (
    // "fmt"
    // "github.com/klauspost/cpuid"
    // "data"
    // "math"
    // "time"
)

type Task struct{
    TID int
    Rec_from map[int]bool   //TID_rec : Received?
    Send_to []int           //list of TID to send to
    PID int                 //Processor ID
} 

func (t * Task)Fire(){
   
}

func (t * Task)ReadyToCompute()bool{
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
// func (t Task)Init(TID int, recFrom map[int]bool, sendTo []int){
//     t.TID = TID
//     t.Rec_from = recFrom
//     t.Send_to = sendTo
// }

// // Find returns the smallest index i at which x == a[i],
// // or len(a) if there is no such index.
// func Find(a []data.Data, x int) int {
//     for i, n := range a {
//         if x == n.TID {
//             return i
//         }
//     }
//     return len(a)
// }



// //checks chan, if senderCountID >= recCountID, update data
// func (t Task)Receive(c <-chan data.Data){
//   //  if t.DiscernChan(c){
//         for recData := range c {
//              fmt.Println("Recvd:\tTID ", t.TID,  "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", recData.Msg)
                    
//             msg := recData.Msg
//             senderID := recData.TID
//             senderCountID := recData.CountID

//             i := Find(t.DataDepVec, senderID) 
//             if i != len(t.DataDepVec){
//                 if senderCountID == t.DataDepVec[i].CountID{
//                     t.DataDepVec[i].Msg = msg
//                     t.DepCount++
//                     fmt.Println("Recvd:\tTID ", t.TID,  "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.DataDepVec[i].Msg,"Msg ID ", senderCountID)
//                     // if cpuid.CPU.LogicalCPU() % 2 ==0{
//                     //     fmt.Println("even ",cpuid.CPU.LogicalCPU())
//                     // }
//                 }
//             }
//         }
//    // }
// }

// //checks if all dep. met
// func (t Task)ReadyToCompute()bool{
//     return len(t.DataDepVec) == t.DepCount
// }

// //computes and fire
// func (t  Task)Compute(c chan data.Data){
//     if t.ReadyToCompute(){
//         t.DepCount = 0

//         var msg int = (t.TID + 1) *100
//         var TID int = t.TID
//         //var countID int = t.DataOutMsg
//         var countID = 0
//         var dataOut data.Data = data.Data{msg, TID, countID}
//         t.Fire(dataOut, c)
//         fmt.Println("Fired:\tTID ", TID, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", dataOut.Msg, "Msg ID ", dataOut.CountID)
//     }
// }

// //opens chan 
// func (t Task)Fire(dataOut data.Data, c chan data.Data){
//     defer close(c)
//     c<-dataOut
// }

// //returns true if task should listen on the channel, since it contains its dep
// func(t Task)DiscernChan(chanIn chan data.Data)bool{
//     var vecTID = make([]int,4)
//     vecTID = t.TID_DEPTable[t.TID]

//     for TID := range vecTID{
//         if t.TID_ChanTable[TID] == chanIn{
//             fmt.Println("TID ",t.TID, " SenderTID ", TID)
//             return true
//         }
//     }
//     return false
// }


// func (t * Task)Producer(iters int) <-chan data.Data {
//     c := make(chan data.Data)
//     go func() {
//         for i := 0; i < iters; i++ {
//             msg := int(math.Pow(10,float64(t.TID))) + t.DataOutMsg
//             c <- data.Data{msg,t.TID,0}
//             // fmt.Println("Fired:\tTID ", t.TID, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", msg, "Orig Data ",t.DataOutMsg)
//             fmt.Println("TID ", t.TID, "Fired and Converted ", t.DataOutMsg, " to ",msg)
//             time.Sleep(1 * time.Second)
//         }
//         close(c)
//     }()
//     return c
// }

// func (t * Task)Consumer(cin <-chan data.Data) {
//     for recData := range cin {
            
//             msg := recData.Msg
//             senderID := recData.TID
//             senderCountID := recData.CountID


//             i := Find(t.DataDepVec, senderID) 
//             if i != len(t.DataDepVec){
//                 if senderCountID == t.DataDepVec[i].CountID{
//                     t.DataDepVec[i].Msg = msg
//                     t.DepCount++
//                     t.DataOutMsg += msg
//                     // fmt.Println("Recvd:\tTID ", t.TID,  "LogicalCPU ", cpuid.CPU.LogicalCPU(), "NewData ", t.DataOutMsg,"Msg ID ", senderCountID)
//                     fmt.Println("TID ", t.TID, " Received  ",t.DataDepVec[i].Msg, " from TID ",senderID)
//                 }
//             }


//         // fmt.Println("Msg ",recData.Msg, "SenderTID ",recData.TID, "MsgCount ",recData.CountID)
//     }


// }

// func (t Task)FanOutUnbuffered(ch <-chan data.Data, size int) []chan data.Data {
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

// func PrintTask(t Task){
//     //fmt.Println("TID ", t.TID, "CountID ",t.DataOutMsg, "LogicalCPU ",cpuid.CPU.LogicalCPU())
// }