package backoff

import (
	"fmt"
	"time"
)

func ExampleMaxDelay() {
	m := MaxDelay(time.Minute, Exponential(time.Second, 5))
	fmt.Println(m)
	for m.Next() {
		fmt.Println(m.Delay())
	}
	//Output:
	//MaxDelay(1m0s, Exp(1s, 5))
	//1s
	//5s
	//25s
	//1m0s
}
