package main

import (
	"fmt"

	ndbc "github.com/clairBuoyant/noaa"
)

func main() {
	fmt.Println(ndbc.GetRealtime("44065"))
}
