package main

import "fmt"

func main() {
	fmt.Println("Learning about arrays in golang")
	// 1. Using Index (for i := 0; i < len(numbers); i++)
	// Use this when you need the position (index) of each element.
	// Scenarios:
	// You want to print both the index and value:
	// fmt.Printf("Index %d, Value %d\n", i, numbers[i])
	// You need to update elements at specific positions.
	// You want to access neighboring elements (e.g., numbers[i+1]).

	// 	2. Using Range (for _, v := range numbers)
	// 	Use this when you only care about the values, not their positions.
	// 	Scenarios:

	// 	Summing all values in the array.
	// 	Printing all values without needing their index.
	// 	Searching for a value (if you don’t need its position).
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[2] = "Banana"
	fmt.Println("Fruits are :", fruits)
	fmt.Println("length od array is : ", len(fruits))
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	var numbers = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Numbers in array  are :", numbers[i])
	}
	//Alternate way to access value
	for _, v := range numbers {
		fmt.Println("Alteranteway: ", v)
	}
	//Array of Strings reverse Print
	var cities = [4]string{"Nellore", "Naidupeta", "Gudur", "Vijayawada"}
	for i := len(cities) - 1; i >= 0; i-- {
		fmt.Println("Cities in reverse", cities[i])
	}
	//Sum of array elements
	var scores = [5]int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum = sum + scores[i]
	}
	fmt.Println("Sum of scores is : ", sum)
	//Alternate way to sum
	var marks = [5]int{10, 20, 30, 40, 50}
	sumOfMarks := 0
	for _, v := range marks {
		sumOfMarks += v
	}
	fmt.Println("Alternate way to sum: ", sumOfMarks)

	var arr = [5]int{1, 2, 3, 4, 5}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("Maximum of arr is : ", max)
	//Given a marks of 5 students in a test fine and print
	//1.Avg Marks
	//The highest marks
	//The lowest Marks

	var namesOfStudents = [5]string{"Thouheedh", "Hafeez", "Ajay", "Sai", "Tez"}
	var marksOfStudents = [5]int{40, 50, 10, 40, 100}
	maxMarks := marksOfStudents[0]
	minMarks := marksOfStudents[0]
	sumMarks := 0
	minName := namesOfStudents[0]
	maxname := namesOfStudents[0]

	fmt.Println("Hi")
	for i := 0; i < len(marksOfStudents); i++ {
		sumMarks = sumMarks + marksOfStudents[i]
		if marksOfStudents[i] > maxMarks {
			maxMarks = marksOfStudents[i]
			maxname = namesOfStudents[i]
		}
		if marksOfStudents[i] < minMarks {
			minMarks = marksOfStudents[i]
			minName = namesOfStudents[i]
		}

	}
	avgMarks := sumMarks / len(marksOfStudents)

	fmt.Println(maxname, maxMarks)
	fmt.Println(minName, minMarks)
	fmt.Println(avgMarks)

	//Check array is sorted

	var sortArray = [5]int{5, 4, 3, 2, 1}
	var isAsc bool = true
	var isDesc bool = true

	for i := 0; i < len(sortArray)-1; i++ {
		if sortArray[i] > sortArray[i+1] {
			isAsc = false
			break
		}
		if sortArray[i] < sortArray[i+1] {
			isDesc = false
		}

	}
	fmt.Println(isAsc)
	fmt.Println(isDesc)
}
