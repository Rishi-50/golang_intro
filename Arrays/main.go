package main

import (
	"fmt"
	"slices"
)

func main()  {

	// ARRAYS 
	// var arr1 [4]int = [4]int{10,5,2,20}
	// fmt.Println(arr1)

	// var arr2 = [3]int{10,20,30}
	// fmt.Println(arr2)

	// objects := [4]string{"bus","bat","chess","dog"}
	// fmt.Println(objects,len(objects))

	// temp := [5]int{1, 2, 3, 4, 5}
    // fmt.Println("dcl:", temp)
	// temp = [...]int{1, 2, 3, 4, 5}
    // fmt.Println("dcl:", temp)
	// temp = [...]int{100, 3: 400, 500}
    // fmt.Println("idx:", temp)

	// var twoD = [2][3]int{
	// 	{1,2,3},
	// 	{10,20,30},
	// }
	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }

	// fmt.Println(twoD)


	// SLICES
	// names := []string{"rishi","ruchit","neel"}
	// fmt.Println(names,len(names))

	// names = append(names,"jinay")

	// fmt.Println(names,len(names))

	// rangeOne := names[0:2]
	// rangeTwo := names[1:]
	// rangeThree := names[:3]

	// fmt.Println(rangeOne,rangeTwo,rangeThree)

	// var s []string 
	// s = make([]string, 3)
	// s[0] = "a"
    // s[1] = "b"
    // s[2] = "c"

    // fmt.Println("set:", s)
    // fmt.Println("get:", s[2])

	// s = append(s, "d")
    // s = append(s, "e", "f")
    // fmt.Println("apd:", s)

	// c := make([]string, len(s))
	// copy(c,s)
	// fmt.Println("cpy:", c)

	// l := s[2:5]
	// fmt.Println("Slice1:",l)

	t1 := []string{"Slice1", "Slice2", "Slice3"}
	// fmt.Println(t1)

	t2 := []string{"Slice1", "Slice2", "Slice"}
	if slices.Equal(t1, t2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
}