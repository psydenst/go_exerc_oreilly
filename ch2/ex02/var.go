package main

import "fmt"
import "math"

func main() {
	var b byte = 255
	var small int32 = math.MaxInt32
	var bigI uint64 = math.MaxInt64

	b += 1
	small += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(small)
	fmt.Println(bigI)

}
