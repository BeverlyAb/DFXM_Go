
package main
// https://golang.org/ref/spec#ImportPath
//https://blog.golang.org/pipelines (multi-listener chan)
// https://flaviocopes.com/golang-event-listeners/
///https://play.golang.org/p/I-GA6DYd0HK //Fanout
import (
    // "github.com/klauspost/cpuid"
    "time"
    // "task"
    // "data"
    "depchantable"
    "fmt"
)



func main(){
    var size int = 19
    var percent int = 50
    dct := new(depchantable.DepChanTable)

    dct.Init(size,percent)
    dct.CreateDAGTable()
    dct.PrintDAGTable()

    dct.CreateTaskSet()
    fmt.Println(dct.TaskSet[0:dct.TaskSize])
    
    runset := createRunSet(size)[:] //slice

    for j := 0; j < 3; j++ {           
        for i := 0; i < size; i++{
            if dct.TaskSet[i].Fire(&dct.TaskSet){
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


//removes task that has ran
//why * isn't needed?
func updateRunSet(runset  []int, index int){
    // runset[index] = runset[len(runset)-1]
    runset[len(runset)-1] = 0
    // runset = runset[:len(runset)-1]
    runset = append(runset[:index],runset[index+1:]...)
    fmt.Println(runset)
}

