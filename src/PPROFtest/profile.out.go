
// https://golang.org/pkg/runtime/pprof/
//https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
// https://echorand.me/posts/golang-dissecting-listen-and-serve/
// https://medium.com/@felipedutratine/profile-your-benchmark-with-pprof-fb7070ee1a94
// https://golang.org/pkg/net/http/pprof/
//https://blog.golang.org/pprof

// http://www.graphviz.org/
package bench
import "testing"
func Fib(n int) int {
    if n < 2 {
      return n
    }
    return Fib(n-1) + Fib(n-2)
}
func BenchmarkFib10(b *testing.B) {
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
      Fib(10)
    }
}