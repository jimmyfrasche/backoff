package backoff

import "fmt"

func ExampleMaxTries() {
	m := MaxTries(3, nil)
	fmt.Println(m)
	for m.Next() {
		fmt.Println(m.Delay())
	}
	//Output:
	//MaxTries(3, None)
	//0s
	//0s
	//0s
}
