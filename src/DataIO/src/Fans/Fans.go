package fans
//pipeline to square numbers
// https://blog.golang.org/pipelines
// https://austburn.me/blog/a-better-fan-in-fan-out-example.html //fans
// https://talks.golang.org/2012/concurrency.slide#47 //timeouts
//https://coderwall.com/p/cp5fya/measuring-execution-time-in-go //elapsed Time
import(
	"fmt"
	"time"
	"log"
)

func main(){
	done := make(chan bool)
	defer close(done)
	var buffer int = 1
	var fanOutSize int = 10
	go FanOut(done,buffer,fanOutSize,1)
	time.Sleep(100*time.Millisecond)	
}

//creates single channel  (SOURCE/SENDER)
func Source(done <-chan bool,nums ...int)<-chan int{
	out := make(chan int)
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
func CopySource(buffer int, src <-chan int, done <- chan bool)[]int{
	out := make([]int,buffer)

	go func(){ 
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
func GenFanOut(buffer int, fanOutSize int)[] chan int{
	out := make([]chan int, fanOutSize) 	
	for i := 0; i < fanOutSize; i++{
		var fanOut = make(chan int, buffer)
		out[i] = fanOut 
	}
	return out
}


//pass data copy to multiple channels (point to points)(FANOUT)
func SplitChannel(dataCopy [] int , chanSet []chan int){
	
	func(){
		for j := 0; j < len(chanSet); j++{
			for i := 0; i < len(dataCopy); i++{
				chanSet[j] <- dataCopy[i]
			}
			close(chanSet[j])
		}
	}()
	
}

//executes fanout channels
func FanOut(done <- chan bool, buffer int, fanOutSize int, nums int){
	defer timeTrack(time.Now(),"FanOut")

	var chanSet []chan int = GenFanOut(buffer, fanOutSize)
	var src_chan <-chan int = Source(done,nums)
	var dataCopy [] int = CopySource(buffer,src_chan,done)
	SplitChannel(dataCopy, chanSet)
	Prints(chanSet)
}
//==========================Time==================================
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

//=========================Prints========================================
//prints all chan set
func Prints(out []chan int){

	for i,_ := range out {
		for n := range out[i]{
			fmt.Println(i, ": ",n)
		} 
	}
}

//prints single channel
func Print(out chan int, index int){
	for n := range out{
		fmt.Println(index, " : ",n)
	}
}






