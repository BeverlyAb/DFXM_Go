
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
    var size int = 5
    var percent int = 50
    dct := new(depchantable.DepChanTable)

    dct.Init(size,percent)
    dct.CreateDAGTable()
    dct.PrintDAGTable()

    dct.CreateTaskSet()
    fmt.Println(dct.TaskSet[0:dct.TaskSize])
    fmt.Println("======================")    
    
    numOfData := 1
    go func (){
        for i := 0; i < size; i++{
            dct.TaskSet[i].Fire(numOfData, dct.TaskSet)
        }
    }()

    time.Sleep(100*time.Millisecond)
    fmt.Println(dct.TaskSet[0:5])

 }