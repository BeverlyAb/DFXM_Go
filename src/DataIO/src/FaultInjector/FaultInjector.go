// package faultinjector
package main
import(
	"fmt"
	"time"
	"math/rand"
	"github.com/klauspost/cpuid"
	"runtime"
)

type FaultInjector struct{
	ReComputeTID map[int]time.Duration 	// TID : timeToSleep
	FaultPID map[int]int 				// PID: uses until failure
}

func (f * FaultInjector)Init(){
	f.ReComputeTID = make(map[int] time.Duration)
	f.FaultPID = make(map[int]int)
}  

//creates map of TID that have computations over the defaultTimeout
func (f * FaultInjector)SetReComputeList(totalTaskSize int, recompSize int, defaultTimeout time.Duration){
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
func (f * FaultInjector)InsertRecompute(tid int){
	timeout, ok := f.ReComputeTID[tid]
	if ok {
		time.Sleep(timeout)
	} 
}


//creates map of PID and the number of uses until failure
//ratio sets upperbound and lowerbound of iterations
func (f * FaultInjector)SetFaultList(totalTaskSize int, cpuSize int, ratio int){
	fmt.Println(runtime.NumCPU(),cpuid.CPU.LogicalCores,cpuid.CPU.LogicalCPU())
	maxUses := totalTaskSize * ratio //uses after failure can be up to ratio x task size
	minUses := totalTaskSize / ratio //unlucky that some tasks fail 

	for i := 0; i < cpuSize; i++ {
		rand.Seed(time.Now().UTC().UnixNano())			//not fixed random
		f.FaultPID[i] = rand.Intn(maxUses-minUses)+minUses
	}

	fmt.Println(f.FaultPID)
}
func (f * FaultInjector)InsertFault(tid int){

}
func main(){
	var totatTaskSize int = 10
	var recompSize int = 5
	var defaultTimeout = time.Millisecond * 50
	var cpuSize int = cpuid.CPU.LogicalCores
	var ratioToFailure int = 3
	faultinjector := new(FaultInjector)
	faultinjector.Init()
	faultinjector.SetReComputeList(totatTaskSize,recompSize,defaultTimeout)
	faultinjector.SetFaultList(totatTaskSize,cpuSize,ratioToFailure)
}