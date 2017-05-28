package main

import "fmt"

func main() {

	// Using range over a "slice/array"
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num

	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i, " with value: ", num)
		}
	}

	// Using range over a "map"
	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
