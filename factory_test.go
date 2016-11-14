package backoff

import "time"

func ExampleFactory() {
	//In a config struct or field
	type Config struct {
		timeout Factory
	}
	c := &Config{}

	//later on, when you need a back off strategy,
	//even if the user did not provide one.
	//In this case, the default of [1s 1s 1s] is used.
	_ = c.timeout.NewOrDefault(func() Strategy {
		return MaxTries(3, Constant(time.Second))
	})
}
