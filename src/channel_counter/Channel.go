package main
//pipeline to square numbers
// https://blog.golang.org/pipelines
// https://austburn.me/blog/a-better-fan-in-fan-out-example.html //fans
// https://talks.golang.org/2012/concurrency.slide#47 //timeouts
import(
	"fmt"
)

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

	func(){ 
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
	

	for j := 0; j < len(chanSet); j++{
		for i := 0; i < len(dataCopy); i++{
			chanSet[j] <- dataCopy[i]
		}
		close(chanSet[j])
	}
	
}

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

//executes fanout channels
func FanOut(done <- chan bool, buffer int, fanOutSize int, nums int){
	var chanSet []chan int = GenFanOut(buffer, fanOutSize)
	var src_chan <-chan int = Source(done,nums)
	var dataCopy [] int = CopySource(buffer,src_chan,done)
	SplitChannel(dataCopy, chanSet)
	Prints(chanSet)
}

func main(){
	done := make(chan bool)
	defer close(done)


	var buffer int = 1
	var fanOutSize int = 10
	FanOut(done,buffer,fanOutSize,1)	
}

/*The merge function converts a list of channels to a single channel by 
starting a goroutine for each inbound channel that copies the values to 
the sole outbound channel. Once all the output goroutines have been started, 
merge starts one more goroutine to close the outbound channel after all sends 
on that channel are done.

Sends on a closed channel panic, so it's important to ensure all sends are 
done before calling close. The sync.WaitGroup type provides a simple way to 
arrange this synchronization:*/
// func merge(cs ...<-chan int) <-chan int {
//     var wg sync.WaitGroup
//     out := make(chan int, 1) // enough space for the unread inputs

//     // Start an output goroutine for each input channel in cs.  output
//     // copies values from c to out until c is closed, then calls wg.Done.
//     output := func(c <-chan int) {
//         for n := range c {
//             out <- n
//         }
//         wg.Done()
//     }
//     wg.Add(len(cs))
//     for _, c := range cs {
//         go output(c)
//     }

//     // Start a goroutine to close out once all the output goroutines are
//     // done.  This must start after the wg.Add call.
//     go func() {
//         wg.Wait()
//         close(out)
//     }()
//     return out
// }

/*stages close their outbound channels when all the send operations are done.
stages keep receiving values from inbound channels until those channels are closed.
This pattern allows each receiving stage to be written as a range loop and 
ensures that all goroutines exit once all values have been successfully 
sent downstream.
*/
