package backoff

import (
	"fmt"
	"testing"
	"time"
)

func TestChain(t *testing.T) {
	c := Chain(OneShot(time.Second), Chain(Fixed(time.Second, time.Second), OneShot(time.Second)))
	var ds []time.Duration
	for c.Next() {
		ds = append(ds, c.Delay())
	}
	if len(ds) != 4 {
		t.Fatal("expected 4 items but got", ds)
	}
	for i, d := range ds {
		if d != time.Second {
			t.Fatalf("expected 1s in %dth position, got %s", i, d)
		}
	}
}

func ExampleChain() {
	c := Chain(Chain(Zero(), Zero()), OneShot(time.Second))
	fmt.Println(c)
	for c.Next() {
		fmt.Println(c.Delay())
	}
	//Output:
	//Chain(Chain(Zero, Zero), OneShot(1s))
	//0s
	//0s
	//1s
}
