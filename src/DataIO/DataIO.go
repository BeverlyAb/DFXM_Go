
package main
// https://golang.org/ref/spec#ImportPath
//https://blog.golang.org/pipelines (multi-listener chan)
// https://flaviocopes.com/golang-event-listeners/
///https://play.golang.org/p/I-GA6DYd0HK //Fanout
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
    // test.PrintTIDTable()
    test.SET_TID_ChanTable()
    // test.PrintChanTable()


    var dataDepVec0 = make([]data.Data,0)

    var dataDepVec123 = make([]data.Data,1)
    dataDepVec123[0] = data.Data{0, 0, 0}

    var dataDepVec4 = make([]data.Data,1)
    dataDepVec4[0] = data.Data{0, 1, 0}

    var dataDepVec5 = make([]data.Data,2)
    dataDepVec5[0] = data.Data{0, 1, 0}
    dataDepVec5[1] = data.Data{0, 2, 0}

    var dataDepVec6 = make([]data.Data,1)
    dataDepVec6[0] = data.Data{0, 2, 0} 

    var dataDepVec7 = make([]data.Data,1)
    dataDepVec7[0] = data.Data{0, 3, 0}   

    var dataDepVec8 = make([]data.Data,4)
    dataDepVec8[0] = data.Data{0, 4, 0}
    dataDepVec8[1] = data.Data{0, 5, 0}
    dataDepVec8[2] = data.Data{0, 6, 0}
    dataDepVec8[3] = data.Data{0, 7, 0}

    t0 :=  task.Task{0, dataDepVec0, 0,cpuid.CPU.LogicalCPU(),0,TID_DEPTable,TID_ChanTable}
    t1 := task.Task{1, dataDepVec123,0, cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t2 := task.Task{2, dataDepVec123,0, cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t3 := task.Task{3, dataDepVec123,0, cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t4 := task.Task{4, dataDepVec4, 0,cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t5 := task.Task{5, dataDepVec5, 0,cpuid.CPU.LogicalCPU(),2,TID_DEPTable,TID_ChanTable}
    t6 := task.Task{6, dataDepVec6, 0,cpuid.CPU.LogicalCPU(),1,TID_DEPTable,TID_ChanTable}
    t7 := task.Task{7, dataDepVec7, 0,cpuid.CPU.LogicalCPU(),2,TID_DEPTable,TID_ChanTable}
    t8 := task.Task{8, dataDepVec8, 0,cpuid.CPU.LogicalCPU(),4,TID_DEPTable,TID_ChanTable}


    //0 Fires
    c0 := t0.Producer(1)
    c0Dep := 3
    chan0 := t0.FanOutUnbuffered(c0, c0Dep)
    t1.TID_ChanTable[t1.TID] = chan0[0]
    t2.TID_ChanTable[t2.TID] = chan0[1]
    t3.TID_ChanTable[t3.TID] = chan0[2]
    go t1.Consumer(t1.TID_ChanTable[t1.TID])
    go t2.Consumer(t2.TID_ChanTable[t2.TID])
    t3.Consumer(t3.TID_ChanTable[t3.TID])

    //1 Fires
    c1 := t1.Producer(1)
    c1Dep := 2
    chan1 := t1.FanOutUnbuffered(c1,c1Dep)
    t4.TID_ChanTable[t4.TID] = chan1[0]
    t5.TID_ChanTable[t5.TID] = chan1[1]
    go t4.Consumer(t4.TID_ChanTable[t4.TID])
    t5.Consumer(t5.TID_ChanTable[t5.TID])

    //2 Fires
    c2 := t2.Producer(1)
    c2Dep := 2
    chan2 := t2.FanOutUnbuffered(c2,c2Dep)
    t5.TID_ChanTable[t5.TID] = chan2[0]
    t6.TID_ChanTable[t6.TID] = chan2[1]
    go t5.Consumer(t5.TID_ChanTable[t5.TID])
    t6.Consumer(t6.TID_ChanTable[t6.TID])

    //3 Fires
    c3 := t3.Producer(1)
    c3Dep := 1
    chan3 := t3.FanOutUnbuffered(c3,c3Dep)
    t7.TID_ChanTable[t7.TID] = chan3[0]
    go t7.Consumer(t7.TID_ChanTable[t7.TID])

    
    //4 Fires
    c4 := t4.Producer(1)
    c4Dep := 1
    chan4 := t4.FanOutUnbuffered(c4,c4Dep)
    t8.TID_ChanTable[t8.TID] = chan4[0]
    go t8.Consumer(t8.TID_ChanTable[t8.TID])

    //5 Fires
    c5 := t5.Producer(1)
    c5Dep := 1
    chan5 := t5.FanOutUnbuffered(c5,c5Dep)
    t8.TID_ChanTable[t8.TID] = chan5[0]
    go t8.Consumer(t8.TID_ChanTable[t8.TID])


    //6 Fires
    c6 := t6.Producer(1)
    c6Dep := 1
    chan6 := t6.FanOutUnbuffered(c6,c6Dep)
    t8.TID_ChanTable[t8.TID] = chan6[0]
    go t8.Consumer(t8.TID_ChanTable[t8.TID])

    //7 Fires
    c7 := t7.Producer(1)
    c7Dep := 1
    chan7 := t7.FanOutUnbuffered(c7,c7Dep)
    t8.TID_ChanTable[t8.TID] = chan7[0]
    go t8.Consumer(t8.TID_ChanTable[t8.TID])    
  
    time.Sleep(100*time.Millisecond)

    // Display 8 contents
     fmt.Println("TID ", t8.TID, "Contains ", t8.DataOutMsg)

}





// func FanOut(ch <-chan data.Data, size, lag int) []chan data.Data {
//     cs := make([]chan data.Data, size)
//     for i, _ := range cs {
//         // The size of the channels buffer controls how far behind the recievers
//         // of the FanOut channels can lag the other channels.
//         cs[i] = make(chan data.Data, lag)
//     }
//     go func() {
//         for i := range ch {
//             for _, c := range cs {
//                 c <- i
//             }
//         }
//         for _, c := range cs {
//             // close all our FanOut channels when the input channel is exhausted.
//             close(c)
//         }
//     }()
//     return cs
// }

