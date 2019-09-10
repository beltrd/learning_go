package main

import (
	"fmt"
	"math"

	"github.com/beltrd/learning_go/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.4))
	fmt.Println(math.Floor(2.5))
	fmt.Println(math.Floor(2.6))
	fmt.Println(math.Ceil(2.6))
	fmt.Println(math.Sqrt(2.6))
	fmt.Println(strutil.Reverse("Diego"))
}
