package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getNumbers(howMany int) []int {
	arr := make([]int, howMany)
	for i := 0; i < howMany; i++ {

		arr[i] = rand.Intn(15)
	}
	return arr
}

func processArr(arrToProcess []int) {
	for _, val := range arrToProcess {
		fmt.Println("Sleep seconds:", val)
		time.Sleep(time.Duration(val) * time.Second)
	}
}

func main() {

	createdArr := getNumbers(5)
	processArr(createdArr)

}
