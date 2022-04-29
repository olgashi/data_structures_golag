package main

import (
	"fmt"
)

func body_temp_sort(temps []float32) (outputSlice []float32) {
  tempCount := make(map[float32]int)
	tempSlice :=  make([]float32, 21)

	for _, num := range temps {
		if _, ok := tempCount[num]; !ok {
			index := int(num * 10) % 97;
			tempSlice[index] = num
		}
		
		tempCount[num]++
	}

	for _, el := range tempSlice {
		if el != 0 {
			if count, ok := tempCount[el]; ok {
				index := 0;
				for index < count {
					outputSlice = append(outputSlice, el)
					index++
				}
			}
		}
	}
  return outputSlice
}

func main() {
	fmt.Println(body_temp_sort([]float32{98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}))
}