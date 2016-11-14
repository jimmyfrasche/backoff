package backoff

import (
	"fmt"
	"time"
)

func ExampleMap() {
	u := MaxDelay(4*time.Second, Linear(0, time.Second))
	fmt.Println(u)

	//Create a minimum delay of 2 seconds
	m := Map(u, func(d time.Duration) time.Duration {
		if d < 2*time.Second {
			return 2 * time.Second
		}
		return d
	})

	for m.Next() {
		fmt.Println(m.Delay())
	}
	//Output:
	//MaxDelay(4s, Linear(0s, 1s))
	//2s
	//2s
	//2s
	//3s
	//4s
}
