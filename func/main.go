package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// вычисляем факториалы от 0 до 10
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d! = %d\r\n", i, factorial(i))
	}
}

// func main() {
// 	may2017 := []int{24, 26, 15, 10, 15, 19, 10, 1, 4, 7, 7, 7, 12, 14, 17,
//         8, 9, 19, 21, 22, 11, 15, 19, 23, 15, 21, 16, 13, 25, 17, 19}
//     may2018 := []int{20, 27, 23, 18, 24, 16, 20, 24, 18, 15, 19, 25, 24, 26,
//         19, 24, 25, 21, 17, 11, 20, 21, 22, 23, 18, 20, 23, 18, 22, 23, 11}

//     comfortCount(may2017)
//     comfortCount(may2018)
// }

// func comfortCount(temperatures []int) {
// 	var count int

// 	for _,v  := range temperatures {
// 		if v >= 22 && v <= 26 {
// 			count++
// 		}
// 	}

// 	fmt.Printf("Количество тёплых дней в этом месяце: %d\n", count)
// }
