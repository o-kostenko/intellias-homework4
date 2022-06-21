package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var (
		result string
		min    = int32(^uint32(0) >> 1)
		max    = -min - 1
	)

	// fmt.Println(min, max)

	input := "7 4325435 -43543253"
	// input := "5"
	// input := "1, 2, 3, 4, 5"
	// input := "1 9 3 4 -5"

	numbers := strings.Split(input, " ")

	for i := range numbers {
		intTmp, err := strconv.Atoi(numbers[i])
		if err != nil {
			log.Fatalln(err)
		}

		int32Tmp := int32(intTmp)

		if int32Tmp > max {
			max = int32Tmp
		}

		if int32Tmp < min {
			min = int32Tmp
		}
	}

	result = fmt.Sprint(max)

	if len(numbers) > 1 {
		result += fmt.Sprint(" ", min)
	}

	fmt.Print(result)
}
