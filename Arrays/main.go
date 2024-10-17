package main

import "fmt"

func main()  {
	// Arrays
	// var arr1 [4]int = [4]int{10,5,2,20}
	// fmt.Println(arr1)

	// var arr2 = [3]int{10,20,30}
	// fmt.Println(arr2)

	// objects := [4]string{"bus","bat","chess","dog"}
	// fmt.Println(objects,len(objects))

	// Slices
	names := []string{"rishi","ruchit","neel"}
	fmt.Println(names,len(names))

	names = append(names,"jinay")

	fmt.Println(names,len(names))

	rangeOne := names[0:2]
	rangeTwo := names[1:]
	rangeThree := names[:3]

	fmt.Println(rangeOne,rangeTwo,rangeThree)
}