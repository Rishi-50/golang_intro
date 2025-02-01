package main

import (
	"fmt"
	// "time"
)

func main()  {

	// count := 5
	// for i := 0; i < count; i++ {
	// 	println("The number is:", i)
	// }

	// if num := 9; num < 0 {
    //     fmt.Println(num, "is negative")
    // } else if num < 10 {
    //     fmt.Println(num, "has 1 digit")
    // } else {
    //     fmt.Println(num, "has multiple digits")
    // }

	// switch time.Now().Weekday() {
	// case time.Monday , time.Wednesday , time.Friday:
	// 	fmt.Println("Today is Odd Day")
	// case time.Thursday , time.Saturday , time.Tuesday:
	// 	fmt.Println("Today is Even Day")
	// default:
	// 	fmt.Println("Today is Sunday")
	// }

	whoami := func(i interface{})  {
		switch t := i.(type) {
		case bool:
			fmt.Println("This is a boolean")
		case string:
			fmt.Println("This is a string")
		case int:
			fmt.Println("This is a int")

		default:
			fmt.Println("Don't know", t)
		}
	}
	whoami(true)
	whoami("Rishi")
	whoami(10)
	whoami(10.352)

}