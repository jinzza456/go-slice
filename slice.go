package main

import "fmt"

func main() {
	// a := make([]int, 2, 3)
	// a[0] = 1
	// a[1] = 2

	// b := append(a, 3)

	// fmt.Printf("%p %p\n", a, b)

	// fmt.Println(a)
	// fmt.Println(b)

	// b[0] = 4
	// b[1] = 5

	// fmt.Println(a)
	// fmt.Println(b)

	// a := []int{1, 2} // a는 2개의 공간만 있음

	// b := append(a, 3) // a에 3을 추가

	// fmt.Printf("%p %p\n", a, b)

	// fmt.Println(a)
	// fmt.Println(b)

	// b[0] = 4
	// b[1] = 5

	// fmt.Println(a)
	// fmt.Println(b)

	// fmt.Println(cap(a))
	// fmt.Println(cap(b))

	// a := make([]int, 2, 4)
	// a[0] = 1
	// a[1] = 2

	// b := make([]int, len(a))

	// for i := 0; i < len(a); i++ {
	// 	b[i] = a[i]
	// }

	// b = append(b, 3)

	// fmt.Printf("%p %p\n", a, b)

	// fmt.Println(a)
	// fmt.Println(b)

	// b[0] = 4
	// b[1] = 5

	// fmt.Println(a)
	// fmt.Println(b)

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// b := a[4:8]
	// fmt.Println(b)

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// b := a[4:]
	// fmt.Println(b)

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// b := a[:4]
	// fmt.Println(b)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 5; i++ {
		var back int
		a, back = RemoveBack(a)
		fmt.Printf("%d, ", back)
	}
	fmt.Println()
	fmt.Println(a)
}

func RemoveBack(a []int) ([]int, int) { // 여러개의 반환값이 나올땐 ()로 묶어준다.
	return a[:len(a)-1], a[len(a)-1]
}
