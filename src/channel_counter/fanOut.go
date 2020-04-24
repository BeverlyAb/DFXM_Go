package main
//pipeline to square numbers
// https://blog.golang.org/pipelines
// https://austburn.me/blog/a-better-fan-in-fan-out-example.html //fans
// https://talks.golang.org/2012/concurrency.slide#47 //timeouts
import(
	"fmt"
)


//Stage 1 (SOURCE): send list of integers on channel. Closes
// INBOUND channel when all int. are sent
func gen(nums ...int)<-chan int{
	out := make(chan int, len(nums)) //buffered
	go func(){
		for _, n := range nums{
			out <- n
		}
		close(out)
	}()
	return out
}

/*Stage 2 (Intermediate): receives integers and returns a
a channel tha emits the square of each received int. After
the inbound chan is closed and all squares are sent, it closes the 
OUTBOUND channel
*/

func square(c <-chan int)<-chan int{
	out := make(chan int)
	go func(){
		for elem := range c{
			square := elem * elem
			out <- square
		}
		close(out)
	}()
	return out
}

/*Stage 3 (SINK) Print's the values from Stage 2 until no more
Sets up pipeline
*/
func main(){
    in := gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := square(in)
    c2 := square(in)

    // Consume the merged output from c1 and c2.
    for n := range merge(c1, c2) {
        fmt.Println(n) // 4 then 9, or 9 then 4
    }
}

/*The merge function converts a list of channels to a single channel by 
starting a goroutine for each inbound channel that copies the values to 
the sole outbound channel. Once all the output goroutines have been started, 
merge starts one more goroutine to close the outbound channel after all sends 
on that channel are done.

Sends on a closed channel panic, so it's important to ensure all sends are 
done before calling close. The sync.WaitGroup type provides a simple way to 
arrange this synchronization:*/
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int, 1) // enough space for the unread inputs

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

/*stages close their outbound channels when all the send operations are done.
stages keep receiving values from inbound channels until those channels are closed.
This pattern allows each receiving stage to be written as a range loop and 
ensures that all goroutines exit once all values have been successfully 
sent downstream.
*/

/*
Need for BUFFER

But in real pipelines, stages don't always receive all the inbound values.
If a stage fails to consume all the inbound values, the goroutines 
attempting to send those values will block indefinitely = Rsrc leak,
Go routines cannot be garbage collected since they contain heap ref.

Need to provide a way for downstream stages to indicate to the senders 
that they will stop accepting input.

*/