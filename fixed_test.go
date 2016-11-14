package backoff

import (
	"fmt"
	"time"
)

func ExampleFixed() {
	f := Fixed(time.Second, time.Minute, time.Hour)
	fmt.Println(f)
	for f.Next() {
		fmt.Println(f.Delay())
	}
	//Output:
	//Fixed(1s, 1m0s, 1h0m0s)
	//1s
	//1m0s
	//1h0m0s
}
