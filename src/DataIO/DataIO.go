
package main
// https://golang.org/ref/spec#ImportPath
//https://blog.golang.org/pipelines (multi-listener chan)
// https://flaviocopes.com/golang-event-listeners/
///https://play.golang.org/p/I-GA6DYd0HK //Fanout
//https://golangbot.com/arrays-and-slices/ //slices
import (
    // "github.com/klauspost/cpuid"
    "time"
    "depchantable"
    "fmt"
    "faultinjector"
)



func main(){
    var totalTaskSize int = 8
    var percent int = 50
    var defaultTimeout = time.Millisecond * 50
    //synthetic problem DAG
    dct := new(depchantable.DepChanTable)
    dct.Init(totalTaskSize,percent,defaultTimeout)
    dct.CreateDAGTable()
    dct.PrintDAGTable()

    //generate Task Set
    dct.CreateTaskSet()
    
    //Creating Tasks that compute over the timeout
    var recompSize int = 5
    faultinjector := new(faultinjector.FaultInjector)
    faultinjector.Init()
    faultinjector.SetReComputeList(totalTaskSize,recompSize,defaultTimeout)

    runset := createRunSet(totalTaskSize) //slice

    for keepRunning(runset){    
        for i := 0; i < totalTaskSize; i++{
            if isInRunSet(runset,i){
                if dct.TaskSet[i].Fire(&dct.TaskSet,faultinjector) {
                    updateRunSet(runset,i)
                }
            }
        }
    }           
    fmt.Println(dct.TaskSet[0:dct.TaskSize])

    time.Sleep(100*time.Millisecond)
 }

//returns arrays of tasks that haven't fired
//only works with no refires for now
 func createRunSet(size int)[]int{
    out := make([]int, size)
    for i:= 0; i < size; i++{
        out[i] = i
    }
    return out
 }


//marks -1 for tasks that ran (why doesn't append truncate)
func updateRunSet(runset  []int, index int){
    runset[index] = -1
    //runset = append(runset[:index],runset[index+1:]...)
}

//checks so that only tasks that haven't ran will run
func isInRunSet(runset[] int,tid int)bool{
    for _,elem := range runset{
        if elem == tid {
            return true
        }
    }
    return false
}

//keep running until all tasks fired (success)
func keepRunning(runset []int)bool{
    for _,tid := range runset{
        if tid != -1{
            return true
        }
    }
    return false
}

