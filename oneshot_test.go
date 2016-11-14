package backoff

import (
	"fmt"
	"time"
)

func ExampleOneShot() {
	o := OneShot(time.Second)
	fmt.Println(o)
	fmt.Println(o.Next())
	fmt.Println(o.Delay())
	fmt.Println(o.Next())
	//Output:
	//OneShot(1s)
	//true
	//1s
	//false
}
