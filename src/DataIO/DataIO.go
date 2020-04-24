
package main
// https://golang.org/ref/spec#ImportPath
//https://blog.golang.org/pipelines (multi-listener chan)
// https://flaviocopes.com/golang-event-listeners/
///https://play.golang.org/p/I-GA6DYd0HK //Fanout
//https://golangbot.com/arrays-and-slices/ //slices
import (
    // "github.com/klauspost/cpuid"
    "time"
    // "task"
    // "data"
    "depchantable"
    "fmt"
)



func main(){
    var size int = 8
    var percent int = 50
    dct := new(depchantable.DepChanTable)

    dct.Init(size,percent)
    dct.CreateDAGTable()
    dct.PrintDAGTable()

    dct.CreateTaskSet()
    //fmt.Println(dct.TaskSet[0:dct.TaskSize])
    
    runset := createRunSet(size) //slice

    for keepRunning(runset){    
        for i := 0; i < size; i++{
            if dct.TaskSet[i].Fire(&dct.TaskSet) && isInRunSet(runset,i){
                updateRunSet(runset,i)
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

func isInRunSet(runset[] int,tid int)bool{
    for _,elem := range runset{
        if elem == tid {
            return true
        }
    }
    return false
}

func keepRunning(runset []int)bool{
    for _,tid := range runset{
        if tid != -1{
            return true
        }
    }
    return false
}

