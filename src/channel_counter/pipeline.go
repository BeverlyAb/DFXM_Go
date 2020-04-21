package main
//pipeline to square numbers
import(
	"fmt"
)


//Stage 1 (SOURCE): send list of integers on channel. Closes
// INBOUND channel when all int. are sent
func gen(nums ...int)<-chan int{
	out := make(chan int)
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
OUTBOUNT channel
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
	c := gen(2,3) 	//Stage 1
	out := square(c)	//Stage 2

	//Consume the outputs
	for n := range out{
		fmt.Println(n)
	} 
}