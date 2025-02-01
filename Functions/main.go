package main

import "fmt"


// func add(a int, b int) (int) {						// normal function
// 	return a + b
// }


// func vals() (int, int, int, int, int) {				// multiple return values function
// 	return 1, 3, 5, 7, 9 
// }


// func sum (nums ...int) {								// variadic function
// 	fmt.Println(nums, " ")
// 	total := 0

// 	for i := 0; i < len(nums); i++ {                    
// 		total += nums[i]
// 	}

// 	fmt.Println("Total", total) 
// }


// func adder() func (int) int{							// closure function
// 	sum := 0
// 	return func(i int) int {
// 		sum += i
// 		return sum
// 	}
// }

// func intseq() func () int{
// 	i:=0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }	

func factorial(n int) int {								// recursive function
	if n == 0 {
		return 1
	} 
	return n * factorial(n-1)
	
}



func main() {
	// a := 10
	// b := 20

	// fmt.Println(add(a,b))					

	// a,b,c,d,e := vals()
	// fmt.Println(a,b,c,d,e)

	// _, _, c, _, _ := vals
	// fmt.Println(c)

	// sum(1,2)
	// sum(10,20,30,40)

	// nums := []int{1,2,10,20,30,40}
	// sum(nums...)

	// pos , neg := adder(), adder()
	
	// for i:=0; i < 10; i++ {
	// 	fmt.Println(pos(i), neg(-2*i))
	// } 

	// nextint := intseq()
	// fmt.Println(nextint())
	// fmt.Println(nextint())
	// fmt.Println(nextint())

	
	fmt.Println(factorial(7))

	var fibonacci func (n int ) int 						// recursive function using anonymous function

	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}

		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println("Fibonacci:", fibonacci(7))
}
