package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learning about slices in go lang")
	// A slice is a flexible, dynamic view into an array.
	// Unlike arrays, slices can grow and shrink in size.
	// Slices are used much more often than arrays in Go.
	var marks = []int{1, 2, 3}
	fmt.Println(marks)
	marks = append(marks, 2)
	fmt.Println(marks)
	//Reverse a slice
	var names = []string{"Thouheedh", "Hafeez", "Tez"}
	rev := []string{}
	for i := len(names) - 1; i >= 0; i-- {
		rev = append(rev, names[i])
	}
	fmt.Println(rev)
	//Reverse a string in place without using extraslice
	var students = []string{"Thouheedh", "Hafeez", "Vikas", "Sai"}
	for i, j := 0, len(students)-1; i < j; i, j = i+1, j-1 {
		students[i], students[j] = students[j], students[i]
	}
	fmt.Println(students)
	//Remove Duplicates from an slice
	var duplicate = []int{1, 2, 2, 2, 2, 3, 3, 4, 4}
	for i := 0; i < len(duplicate)-1; i++ {
		if duplicate[i] == duplicate[i+1] {
			duplicate = append(duplicate[:i], duplicate[i+1:]...)
			i--
			//if i-- is not used it will fail to recheck at index 1
		}
	}
	fmt.Println(duplicate)
	//Find Max and Min in a given slice
	var numbers = []int{1, 2, 3, 4, 5}
	minNumbers := numbers[0]
	maxNumbers := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < minNumbers {
			minNumbers = numbers[i]
		}
		if numbers[i] > maxNumbers {
			maxNumbers = numbers[i]
		}
	}
	fmt.Println(minNumbers)
	fmt.Println(maxNumbers)
	//Alternate way for min and max
	var givenNumbers = []int{2, 4, 1, 5, 7, 8}
	minNum := givenNumbers[0]
	maxNum := givenNumbers[0]
	for _, v := range givenNumbers {
		if v < minNum {
			minNum = v
		}
		if v > maxNum {
			maxNum = v
		}
	}
	fmt.Println(minNum)
	fmt.Println(maxNum)

	//Calculate the sum of elements in a given slice
	var num = []int{10, 20, 30}
	sum := 0
	for _, v := range num {
		sum += v
	}
	fmt.Println(sum)
	//Check if a slice contains a value
	var myNames = []string{"Hafeez", "Thouheedh", "Thouhee"}
	myNickname := "Thouhee"
	isExist := false
	for _, v := range myNames {
		if myNickname == v {
			isExist = true
		}
	}
	fmt.Println(isExist)
	//Sort a slice in asce/desc order
	var unSorted = []int{1, 4, 2, 7, 3, 9}
	sort.Ints(unSorted)
	fmt.Println(unSorted)
	nums := []int{1, 2, 3}
	data, _ := json.Marshal(nums) // converts slice to JSON
	fmt.Println(string(data))

	//Learning how to use "make" in slices
	//s := make([]Type, length, capacity)
	usemake := make([]int, 0, 5)
	//Slice has 2 elements, but can grow up to capacity 5 before reallocating memory.
	fmt.Println(usemake)
	for i := 0; i < cap(usemake); i++ {
		usemake = append(usemake, i)

	}
	fmt.Println(usemake)

	//Append two slices
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7}
	merged := append(a, b...)
	fmt.Println(merged)
	for i := 0; i < len(merged); i++ {
		fmt.Println(merged[i])
	}
	//Delete an element at specific index
	deleteElement := []int{1, 2, 3, 4, 5}
	cleanedEle := append(deleteElement[:3], deleteElement[4:]...)
	fmt.Println(cleanedEle)
	//Copy one slice data to other slice
	src := []int{4, 6, 7, 8, 3, 5}
	dest := make([]int, len(src))
	copy(dest, src)
	fmt.Println(dest)

}
