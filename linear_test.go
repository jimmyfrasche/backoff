package backoff

import (
	"fmt"
	"time"
)

func ExampleLinear() {
	lb := MaxTries(3, Linear(4*time.Second, 5*time.Second))
	fmt.Println(lb)
	for lb.Next() {
		fmt.Println(lb.Delay())
	}
	//Output:
	//MaxTries(3, Linear(4s, 5s))
	//4s
	//9s
	//14s
}
