
package main
// https://golang.org/ref/spec#ImportPath
//https://blog.golang.org/pipelines (multi-listener chan)
// https://flaviocopes.com/golang-event-listeners/
///https://play.golang.org/p/I-GA6DYd0HK //fanout
import (
    "github.com/klauspost/cpuid"
    "time"
    "task"
    "data"
    "depchantable"
    "fmt"
)

//creates data dep. of up to size, dep. are from T0 to T(size) -1
func CreateDepVec(size int)[]data.Data{
    var dataDepVec = make([]data.Data,size)
    for i  := 0; i < size; i++ {
        dataDepVec[i] = data.Data{0, i, 0}
    }
    return dataDepVec
}

func main(){
    TID_DEPTable := make(map[int][]int)
    TID_ChanTable := make(map[int]chan data.Data)
    test := depchantable.DepChanTable{TID_DEPTable,TID_ChanTable}
    test.SET_TID_DEPTable()
    test.PrintTIDTable()
    test.SET_TID_ChanTable()
    test.PrintChanTable()


    var dataDepVec0 = make([]data.Data,0)

    var dataDepVec123 = make([]data.Data,1)
    dataDepVec123[0] = data.Data{0, 0, 0}

    var dataDepVec4 = make([]data.Data,1)
    dataDepVec4[0] = data.Data{0, 1, 0}

    // var dataDepVec5 = make([]data.Data,2)
    // dataDepVec5[0] = data.Data{0, 1, 0}
    // dataDepVec5[1] = data.Data{0, 2, 0}

    // var dataDepVec6 = make([]data.Data,1)
    // dataDepVec6[0] = data.Data{0, 2, 0} 

    // var dataDepVec7 = make([]data.Data,1)
    // dataDepVec7[0] = data.Data{0, 3, 0}   

    // var dataDepVec8 = make([]data.Data,4)
    // dataDepVec8[0] = data.Data{0, 4, 0}
    // dataDepVec8[1] = data.Data{0, 5, 0}
    // dataDepVec8[2] = data.Data{0, 6, 0}
    // dataDepVec8[3] = data.Data{0, 7, 0}

    t0 :=  task.Task{0, dataDepVec0, cpuid.CPU.LogicalCPU(),0,TID_DEPTable,TID_ChanTable}
    t1 := task.Task{1, dataDepVec123, cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t2 := task.Task{2, dataDepVec123, cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t3 := task.Task{3, dataDepVec123, cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    // t4 := task.Task{4, dataDepVec4, cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    // t5 := task.Task{5, CreateDepVec(2), cpuid.CPU.LogicalCPU(),2,TID_DEPTable,TID_ChanTable}
    // t6 := task.Task{6, CreateDepVec(2), cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    // t7 := task.Task{7, CreateDepVec(2), cpuid.CPU.LogicalCPU(),2,TID_DEPTable,TID_ChanTable}
    // t8 := task.Task{8, CreateDepVec(2), cpuid.CPU.LogicalCPU(),4,TID_DEPTable,TID_ChanTable}




    // c0 := make(chan data.Data)
    // c1 := make(chan data.Data)

    // go t0.Compute(test.TID_ChanTable[0])
    
    // go t1.Receive(test.TID_ChanTable[0])
    // go t2.Receive(test.TID_ChanTable[0])
    // go t3.Receive(test.TID_ChanTable[0])

    // go t1.Compute(test.TID_ChanTable[1])
    // go t2.Compute(test.TID_ChanTable[2])
    // go t3.Compute(test.TID_ChanTable[3])

    // go t4.Receive(test.TID_ChanTable[1])
    // go t4.Receive(test.TID_ChanTable[2])

    c := producer(1,t0)
    // chans := fanOutUnbuffered(c, 3)
    // go consumer(chans[0],t1)
    // go consumer(chans[1],t2)
    // consumer(chans[2],t3)

    chans := fanOutUnbuffered(c, 3)
    t1.TID_ChanTable[t1.TID] = chans[0]
    t2.TID_ChanTable[t2.TID] = chans[1]
    t3.TID_ChanTable[t3.TID] = chans[2]
    go consumer(chans[0],t1)
    go consumer(chans[1],t2)
    consumer(chans[2],t3)

    time.Sleep(100*time.Millisecond)
}

func producer(iters int, t task.Task) <-chan data.Data {
    c := make(chan data.Data)
    go func() {
        for i := 0; i < iters; i++ {
            c <- data.Data{(t.TID+1)*100,t.TID,0}
            fmt.Println("Fired:\tTID ", t.TID, "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", (t.TID+1)*100, "Msg ID ", 0)
            time.Sleep(1 * time.Second)
        }
        close(c)
    }()
    return c
}

func consumer(cin <-chan data.Data, t task.Task) {
    for recData := range cin {
            
            msg := recData.Msg
            senderID := recData.TID
            senderCountID := recData.CountID

            i := task.Find(t.DataDepVec, senderID) 
            if i != len(t.DataDepVec){
                if senderCountID == t.DataDepVec[i].CountID{
                    t.DataDepVec[i].Msg = msg
                    t.DepCount++
                    fmt.Println("Recvd:\tTID ", t.TID,  "LogicalCPU ", cpuid.CPU.LogicalCPU(), "Data ", t.DataDepVec[i].Msg,"Msg ID ", senderCountID)
                }
            }


        // fmt.Println("Msg ",recData.Msg, "SenderTID ",recData.TID, "MsgCount ",recData.CountID)
    }


}

// func fanOut(ch <-chan data.Data, size, lag int) []chan data.Data {
//     cs := make([]chan data.Data, size)
//     for i, _ := range cs {
//         // The size of the channels buffer controls how far behind the recievers
//         // of the fanOut channels can lag the other channels.
//         cs[i] = make(chan data.Data, lag)
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

func fanOutUnbuffered(ch <-chan data.Data, size int) []chan data.Data {
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