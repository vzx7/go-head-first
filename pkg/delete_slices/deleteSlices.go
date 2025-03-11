package deleteslices

import (
	"fmt"
	"os"
	"strconv"
)

func Del(method int) {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using method", method)
	fmt.Println("Using index", i)
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	if method == 1 {
		aSlice = append(aSlice[:i], aSlice[i+1:]...)
		fmt.Println("After 1st deletion:", aSlice)
	} else if method == 2 {
		aSlice[i] = aSlice[len(aSlice)-1]
		aSlice = aSlice[:len(aSlice)-1]
		fmt.Println("After 2st deletion:", aSlice)
	}
}
