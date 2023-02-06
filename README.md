# raffle

## usage

```go
package main

import (
	"fmt"
	"github.com/dlwm/raffle"
)

func main() {
	var elList []raffle.Elem
	for i := 0; i < 10; i++ {
		elList = append(elList, raffle.NewElems(99, uint32(i+1), 1, i+1)...)
	}
	p := raffle.NewPool(elList)

	_, cList := p.RepeatDraw(10000)
	resList := make([]int, 11)
	for _, content := range cList {
		resList[content.(int)]++
	}

	fmt.Println(resList)
}

```