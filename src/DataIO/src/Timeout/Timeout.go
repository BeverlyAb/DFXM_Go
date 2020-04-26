package timeout
//Used to handle Task re-computes
import(
	// "fmt"
	"time"
)

type Timeout struct{
	Start time.Time
	Elapsed time.Duration
	HasTimedOut bool
	TaskTimeout time.Duration
}

func (tm * Timeout)Init(taskTimeout time.Duration){
	tm.Start = time.Now()
    tm.Elapsed = time.Now().Sub(tm.Start)
    tm.TaskTimeout = taskTimeout
    tm.HasTimedOut = tm.Elapsed > tm.TaskTimeout
}

//updates elapsed time and checks if computation has timed out
func (tm* Timeout)Update(){
	tm.Elapsed = time.Now().Sub(tm.Start)
    tm.HasTimedOut = tm.Elapsed > tm.TaskTimeout 
}