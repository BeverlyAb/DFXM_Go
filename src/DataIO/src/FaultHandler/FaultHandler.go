// package faulthandler
package main
import(
	"fmt"
	"time"
	"math/rand"
)

type FaultHandler struct{
	ReComputeTID map[int]time.Duration // TID : timeToSleep
}

func (f * FaultHandler)Init(){
	f.ReComputeTID = make(map[int] time.Duration)
}  

//creates map of TID that have computations over the defaultTimeout
func (f * FaultHandler)SetReComputeList(totalTaskSize int, recompSize int, defaultTimeout time.Duration){
	timeLimitOver := int(defaultTimeout * 3) //compute time can be over 3x task timeout period
	minTimeLimit := int(defaultTimeout) + 1
	 
	for len(f.ReComputeTID) < recompSize {
		rand.Seed(time.Now().UTC().UnixNano())			//not fixed random	
		recomputeTID := rand.Intn(100) % totalTaskSize 
		f.ReComputeTID[recomputeTID] = time.Duration(rand.Intn(timeLimitOver-minTimeLimit)+minTimeLimit)
	}

	fmt.Println(f.ReComputeTID)
}

//creates a simulated long computation time if the TID is within ReComputeTID
func (f * FaultHandler)InsertRecompute(tid int){
	timeout, ok := f.ReComputeTID[tid]
	if ok {
		time.Sleep(timeout)
	} 
}

func main(){
	var totatTaskSize int = 10
	var recompSize int = 5
	var defaultTimeout = time.Millisecond * 50

	faulthandler := new(FaultHandler)
	faulthandler.Init()
	faulthandler.SetReComputeList(totatTaskSize,recompSize,defaultTimeout)
}