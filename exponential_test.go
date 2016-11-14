package backoff

import (
	"fmt"
	"time"
)

func ExampleExponential() {
	e := ZeroThen(MaxTries(4, Exponential(time.Second, 2)))
	fmt.Println(e)
	for e.Next() {
		fmt.Println(e.Delay())
	}
	//Output:
	//Chain(Zero, MaxTries(4, Exp(1s, 2)))
	//0s
	//1s
	//2s
	//4s
	//8s
}
