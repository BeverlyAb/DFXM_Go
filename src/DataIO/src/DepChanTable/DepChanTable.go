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

func (dct * DepChanTable)Init(size int, percent int){
	dct.TaskSize = size
	dct.percentageOfCon = percent
}

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
	// fmt.Println(dct.DAGTable)
	dct.PrintDAGTable()
	// fmt.Println(dct.DAGTable)
	// TID_DEPTable := make(map[int][]int)
	// TID_ChanTable := make(map[int]chan data.Data)
	// test := DepChanTable{TID_DEPTable,TID_ChanTable}
	// test.SET_TID_DEPTable()
	// fmt.Println("8th's first ",test.TID_DEPTable[8][0])
	// test.SET_TID_ChanTable()
	// fmt.Println("Channel 5 ",test.TID_ChanTable[5])
}