//package depchantable
 package main
//https://www.golangprograms.com/constructors-in-golang.html
import (
   "fmt"
   // "data"
   "task"
   "math/rand"
)


type DepChanTable struct{
	TaskSize int
	TaskSet [] task.Task
	percentageOfCon int
	DAGTable [100][100] int //rec x send
} 

//init, percentageOfCon = % of forward directed edges
func (dct * DepChanTable)Init(size int, percent int){
	dct.TaskSize = size
	dct.percentageOfCon = percent
}

//Creates a pseudo-random DAG 
func (dct * DepChanTable)CreateDAGTable(){
	for i := 0; i < dct.TaskSize; i++ {
	    for j := i+1; j < dct.TaskSize; j++  {
	    	random := rand.Intn(100) // fixed random
	       	if random < dct.percentageOfCon {
	       		dct.DAGTable[j][i] = 1 //inverted, so that lower left triangle is populated
	       }
	    }
	} 
}

func (dct DepChanTable)CreateTaskSet(){

}

//creates an array of TIDs that the indexed task should send to
func (dct DepChanTable)createSendTo(index int)[]int{
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
func (dct DepChanTable)createRecFrom(index int)map[int]bool{
	recFrom := make(map[int]bool)
	for j := 0; j < dct.TaskSize; j++ {
		if dct.DAGTable[index][j] == 1{
			recFrom[j] = false
		}
	}
	return recFrom
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

func main(){
	var size int = 5
	var percent int = 50
	dct := new(DepChanTable)

	dct.Init(size,percent)
	dct.CreateDAGTable()
	dct.PrintDAGTable()
	fmt.Println(dct.createSendTo(1))
	fmt.Println(dct.createRecFrom(3))
}