package main

import "fmt"

func main() {
	records := make([][]string, 0)

	//student 1
	student1 := make([]string, 5)
	student1[0] = "Jones"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "0.00"
	records = append(records, student1)

	student2 := make([]string, 5)
	student2[0] = "Doe"
	student2[1] = "Jane"
	student2[2] = "100.00"
	student2[3] = "0.00"
	records = append(records, student2)
	fmt.Println(records)

	fmt.Println(records[1])
	fmt.Println(len(records))
}
