# Async package
Schedules generic tasks for asynchronous execution. Tasks that are scheduled with the same group will be executed sequentially, preserving the order of schedule.
Async func return Future

```go

type Future[T any] interface {
	Done() bool
	Value() T
	Wait() T
}

```

```go
package main

import (
	"fmt"
	"time"

	"github.com/antonchirikalov/async"
)

const bufferSize = 10

func main() {
	pool := async.NewPool[string](bufferSize)
	future := pool.Async("group-name", func() string {
		<-time.After(1 * time.Microsecond)
		return "hello"
	})
	res := future.Wait()
	fmt.Println(res)
}


```