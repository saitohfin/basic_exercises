package pointers

import "fmt"

//Bitcoin is a kind of virtual money
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
