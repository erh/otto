package otto

import "fmt"

const LIMIT_STRING = 1000000

func checkString(s string) {
	if len(s) > LIMIT_STRING {
		err := fmt.Errorf("String too big: %d > %d", len(s), LIMIT_STRING)
		panic(err)
	}
}
