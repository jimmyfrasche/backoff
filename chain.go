package backoff

import (
	"fmt"
	"strings"
	"time"
)

type chain struct {
	strategies []Strategy
}

//Chain returns a Strategy that uses each strategy in turn.
//When the first strategy returns Next() == false, it is discarded and the second is used.
//This proceeds until there are no strategies left in the chain.
//
//If no strategies are provided, the chain immediately reports that Next() == false.
//
//A nil Strategy cannot be passed to Chain.
func Chain(strategies ...Strategy) Strategy {
	for _, s := range strategies {
		if s == nil {
			panic("chain cannot process a nil Strategy")
		}
	}
	return &chain{
		strategies: strategies,
	}
}

//ZeroThen with a nil strategy is equivalent to Zero.
//ZeroThen with a nonnil strategy is equivalent to Chain(Zero(), strategy).
func ZeroThen(strategy Strategy) Strategy {
	z := Zero()
	if strategy == nil {
		return z
	}
	return Chain(z, strategy)
}

func (c *chain) Next() bool {
	if len(c.strategies) == 0 {
		return false
	}
	for !c.strategies[0].Next() {
		c.strategies = c.strategies[1:]
		if len(c.strategies) == 0 {
			return false
		}
	}
	return true
}

func (c *chain) Delay() time.Duration {
	if len(c.strategies) == 0 {
		panic(errNoNext)
	}
	return c.strategies[0].Delay()
}

func (c *chain) String() string {
	var items []string
	for _, v := range c.strategies {
		s := "<unknown>"
		st, ok := v.(fmt.Stringer)
		if ok {
			s = st.String()
		}
		items = append(items, s)
	}
	return "Chain(" + strings.Join(items, ", ") + ")"
}
