// package fans
package main
//pipeline to square numbers
// https://blog.golang.org/pipelines
// https://austburn.me/blog/a-better-fan-in-fan-out-example.html //fans
// https://talks.golang.org/2012/concurrency.slide#47 //timeouts
//https://coderwall.com/p/cp5fya/measuring-execution-time-in-go //elapsed Time
import(
	"fmt"
	"time"
	"log"
	"data"
)

func main(){
	done := make(chan bool)
	defer close(done)
	var buffer int = 1
	var fanOutSize int = 10
	msg := 9000
	TID := 0
	CountID := 0
	go Prints(FanOut(done,buffer,fanOutSize,data.Data{msg,TID,CountID}))
	time.Sleep(100*time.Millisecond)	
}

//creates single channel  (SOURCE/SENDER)
func Source(done <-chan bool,nums ...data.Data)<-chan data.Data{
	out := make(chan data.Data)
	go func(){
		for _, n := range nums{
			select{
			case <-done:
				return 
			case out<-n:
			}
		}
		close(out)
	}()
	
	return out
}

//creates a copy of data from source
func CopySource(buffer int, src <-chan data.Data, done <- chan bool)[]data.Data{
	out := make([]data.Data,buffer)

	func(){  //no go
		for i := 0; i < buffer; i++{
			select{
			case <-done:
				return 
			case out[i] = <-src:
			}
		}
	}()
	return out
}

//creates channel with amount of elems on sender channel specified (fan generator)
func GenFanOut(buffer int, fanOutSize int)[] chan data.Data{
	out := make([]chan data.Data, fanOutSize) 	

	for i := 0; i < fanOutSize; i++{
		var fanOut = make(chan data.Data, buffer)
		out[i] = fanOut 
	}
	return out
}


//pass data copy to multiple channels (point to points)(FANOUT)
func SplitChannel(dataCopy [] data.Data , chanSet []chan data.Data){
	
	go func(){
		for j := 0; j < len(chanSet); j++{
			for i := 0; i < len(dataCopy); i++{
				chanSet[j] <- dataCopy[i]
			}
			close(chanSet[j])
		}
	}()
	
}

//executes fanout channels and outputs channels with sent data
func FanOut(done <- chan bool, buffer int, fanOutSize int, nums data.Data)[] chan data.Data{
	defer timeTrack(time.Now(),"FanOut")

	var chanSet []chan data.Data = GenFanOut(buffer, fanOutSize)
	var src_chan <-chan data.Data = Source(done,nums)
	var dataCopy [] data.Data = CopySource(buffer,src_chan,done)
	SplitChannel(dataCopy, chanSet)
	return chanSet
}
//==========================Time==================================
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

//=========================Prints========================================
//prints all chan set
func Prints(out []chan data.Data){

	for i,_ := range out {
		for n := range out[i]{
			fmt.Println(i, ": ",n.Msg)
		} 
	}
}

//prints single channel
func Print(out chan data.Data, index int){
	for n := range out{
		fmt.Println(index, " : ",n.Msg)
	}
}






