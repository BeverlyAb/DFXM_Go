package depchantable
 // package main
//https://www.golangprograms.com/constructors-in-golang.html
import (
   "fmt"
   "task"
   "math/rand"
   "data"
   "time"
)


type DepChanTable struct{ 
	TaskSize int
	TaskSet [] task.Task
	percentageOfCon int
	DAGTable [][] int //rec x send
	Timeout time.Duration
} 

//init, percentageOfCon = % of forward directed edges
func (dct * DepChanTable)Init(size int, percent int, timeout time.Duration){
	dct.TaskSize = size
	dct.percentageOfCon = percent
	dct.Timeout = timeout
	dct.TaskSet = make([]task.Task, dct.TaskSize)
	dct.DAGTable = make([][]int, 0)
}

//Creates a pseudo-random DAG 
func (dct * DepChanTable)CreateDAGTable(){
	
	tmp := make([]int,0)
	dct.DAGTable = make([][]int,0)
	for i := 0; i < dct.TaskSize; i ++{

	      tmp = make([]int, 0)
	      for j := 0; j < dct.TaskSize; j ++{
	            tmp = append(tmp, 0)
	      }
	    dct.DAGTable = append(dct.DAGTable, tmp)
	} 
	
	for i := 0; i < dct.TaskSize; i++ {
	    for j := i+1; j < dct.TaskSize; j++  {
	    	random := rand.Intn(100) // fixed random
	       	if random < dct.percentageOfCon {
	       		dct.DAGTable[j][i] = 1 //inverted, so that lower left triangle is populated	
	       }
	    }
	} 
}

//creates Tasks with personal send and receive 
func (dct * DepChanTable)CreateTaskSet(){
	for i := 0; i < dct.TaskSize; i++ {
		tID := i
		var rec_from map[int] bool = dct.createRecFrom(i)
		var send_to []int = dct.createSendTo(i)
		pID := 0
		var dataRecvd [] data.Data
		var timeout time.Duration = dct.Timeout
		var refireCount int = 0
		dct.TaskSet[i] = task.Task{	tID, 
									rec_from, 
									send_to,pID, 
									dataRecvd, 
									timeout, 
									refireCount}
	}
}

//creates an array of TIDs that the indexed task should send to
func (dct * DepChanTable)createSendTo(index int)[]int{
	sendTo := make([]int,dct.TaskSize)
	k := 0
	for i := 0; i < dct.TaskSize; i++ {
		if dct.DAGTable[i][index] == 1{
			sendTo[k] = i
			k++
		}
	}
	return sendTo[0:k]
}

//creates map of TIDs that the indexed task should receive
//data from. By default all value pair are set to false
//since it has not received data from them
func (dct * DepChanTable)createRecFrom(index int)map[int]bool{
	recFrom := make(map[int]bool)
	for j := 0; j < dct.TaskSize; j++ {
		if dct.DAGTable[index][j] == 1{
			recFrom[j] = false
		}
	}
	return recFrom
}
// func (dct * DepChanTable)RemoveTask(index int){
// 	dct.TaskSet[index].Invalidate()
// }

func(dct * DepChanTable)TasksComplete()bool{
	for i := 0; i < dct.TaskSize -1; i++{//exclude last task
		if dct.TaskSet[i].TID != -1{
			return false
		}
	}
	return true
}


//prints DAG
func (dct DepChanTable)PrintDAGTable(){
	for i := 0; i < dct.TaskSize; i++ {
	    for j := 0; j < dct.TaskSize; j++  {
	    	fmt.Print(dct.DAGTable[i][j], " ")
	    }
	    fmt.Println()
	} 
}

//test
func main(){
	// var size int = 5
	// var percent int = 50
	// dct := new(DepChanTable)

	// dct.Init(size,percent)
	// dct.CreateDAGTable()
	// dct.PrintDAGTable()
	
// 	fmt.Println(dct.createSendTo(1))
// 	fmt.Println(dct.createRecFrom(3))
	
// 	dct.CreateTaskSet()
// 	fmt.Println(dct.TaskSet[0:dct.TaskSize])
}