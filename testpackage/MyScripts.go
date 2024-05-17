package testpackage

import (
	"fmt"
	"strconv"
	"strings"
)

func HelloWorld() {
	fmt.Println("Hello World!")
}

// 0
func Quiz0() {
	//จากหน้า 22 จงเปลี่ยนให้เป็น if
	// i := 2
	// fmt.Println("Example-: Switch case condition")
	// switch i {
	// case 0:
	// 	fmt.Println("Zero")
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Your i not in case.")
	// }
	i := 2
	fmt.Println("Example-: If-Else condition")
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Your i not in case.")
	}
	fmt.Println("----------------------")
}

// 1
func Quiz1() {
	//ระหว่าง 1-100 มีเลขที่หาร3ลงตัวกี่ตัว อะไรบ้าง (for if)
	max := 100
	count := 0

	for i := 1; i <= max; i++ {
		if i%3 == 0 {
			count++
		}
	}
	arr := make([]int, count)
	index := 0
	max = 100

	for i := 1; i <= max; i++ {
		if i%3 == 0 {
			arr[index] = i
			index++
		}
	}

	fmt.Printf("มีทั้งหมด : %d ตัว ได้แก่\n", count)
	fmt.Printf("\n{\t")
	for i := 0; i < count; i++ {
		fmt.Printf("%d, ", arr[i])
	}
	fmt.Printf("\t}\n\n")
	fmt.Println("----------------------")
}

// 1_2
func num(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func Quiz1_2() {
	//สร้างฟังชั่นคำนวณเลขยกกำลัง เช่น num(20,2) (เรียกใช้เป็นฟังชั่น) ฝึกการใช้ฟังชั่น (หน้า34-35)
	a := 20
	b := 4
	result := num(a, b)
	fmt.Printf("\nResult: %d + %d = %d\n---------------\n", a, b, result)
}

// 2

// func MinSort(arr []int) []int {
// 	n := len(arr)
// 	for i := 0; i < n; i++ {
// 		swapped := false
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 				swapped = true
// 			}
// 		}

// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// func MaxSort(arr []int) []int {
// 	n := len(arr)
// 	for i := 0; i < n; i++ {
// 		swapped := false
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] < arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 				swapped = true
// 			}
// 		}

//			if !swapped {
//				break
//			}
//		}
//		return arr
//	}
func MinSort(arr []int) int {
	var max int
	for _, e := range arr {
		if e < max || max == 0 {
			max = e
		}
	}
	return max
}
func MaxSort(arr []int) int {
	var min int
	for _, e := range arr {
		if e > min || min == 0 {
			min = e
		}
	}
	return min
}

func Quiz2() {
	//เขียนโปรแกรมหาเลขที่น้อยสุดและมากสุด (เรียกใช้เป็นฟังชั่น) ฝึกการใช้ฟังชั่น (หน้า34-35)
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Printf("Array count : %d\n", len(x))
	fmt.Println("Original array:", x)
	fmt.Println("Min number is : ", MinSort(x))
	// fmt.Println("Min Sorted array:  ", x)
	fmt.Println("Max number is : ", MaxSort(x))
	// fmt.Println("Max Sorted array:  ", x)
	fmt.Println("--------")
}

// 3
func Quiz3() {
	// 1-1000 มีเลข9รวมทั้งหมดกี่ตัว จงเขียนโปรแกรม(คำตอบข้อ3 300)
	count := 0
	for i := 1; i <= 1000; i++ {
		t := strconv.Itoa(i)
		count += strings.Count(t, "9")
	}
	fmt.Println("1-1000 number 9")
	fmt.Println("Result: ", count)
}

// 3_1
func someFunc(number int) {
	// 1-1000 มีเลข9รวมทั้งหมดกี่ตัว จงเขียนโปรแกรม(คำตอบข้อ3 300)
	count := 0
	for i := 1; i <= number; i++ {
		t := strconv.Itoa(i)
		count += strings.Count(t, "9")
	}
	fmt.Println(number, " : find number 9")
	fmt.Println("Result: ", count)
}

func Quiz3_1() {
	//(เรียกใช้เป็นฟังชั่น) ใส่ค่าตัวเลขเข้าไปให้ฟังชั่น เพื่อหาเลข9ในจำนวนเลขที่ใส่เข้าไป เช่น someFunc(10000) หาเลข9ตั้งแต่1-10000
	myNum := 10000
	someFunc(myNum)
}

// 4
func Quiz4() {
	//var myWords = "AW SOME GO!" ตัดช่องว่างออกโดยใช้ forloop “AWSOMEGO!”
	var myWords = "AW SOME GO!"
	var result string

	for _, char := range myWords {
		if char != ' ' {
			result += string(char)
		}
	}

	fmt.Println("result: ", result)

}

// 4_1
func cutText(values string) {
	for _, v := range values {
		if string(v) == " " {
			strings.TrimSpace(string(v))
		} else {
			fmt.Printf("%s", string(v))
		}
	}
}

func Quiz4_1() {
	// ทำเป็นฟังชัน cutText(“ine t”) ⇒ “inet”
	character := "i n et"
	cutText(character)
}

// 5
func Maps2() {
	var maxemp int
	var fname, lname string

	peoples := make(map[int]map[string]string)
	fmt.Printf("How many Employee register : ")
	fmt.Scanf("%d", &maxemp)
	for i := 0; i < maxemp; i++ {
		fmt.Println("employee : ", i+1)
		fmt.Printf("First Name: ")
		fmt.Scan(&fname)
		fmt.Printf("Last Name: ")
		fmt.Scan(&lname)
		peoples[i+1] = map[string]string{
			"fname": fname,
			"lname": lname,
		}
	}

	fmt.Println("Example-: Maps")
	fmt.Println(peoples)
	for i := 0; i < maxemp; i++ {
		fmt.Printf("%d: %s %s\n", i, peoples[i+1]["fname"], peoples[i+1]["lname"])
	}
}

func Quiz5() {
	// peoples := map[string]map[string]string{
	// 	"emp_01": {
	// 		"fname": "Marry",
	// 		"lname": "Jane",
	// 	},
	// 	"emp_02": {
	// 		"fname": "Gwenn",
	// 		"lname": "Steframe",
	// 	},
	// }

	Maps2()
}

// (ดูหน้า29เป็นตัวอย่าง) เรื่อง map

// 6
type mycompany struct {
	role   string
	name   string
	salary float64
}

func Quiz6() {
	// สร้างStruct company เก็บโครงสร้างข้อมูลบริษัท ให้ออกแบบเองว่าจะมีอะไรในstructบ้าง -ลองแอดค่า และprintมาโชว์ (หน้า45)

	var emp mycompany
	emp1 := new(mycompany)
	emp.role = "CEO"
	emp.name = "pab"
	emp.salary = 60000.00

	emp1.role = "Director"
	emp1.name = "tae"
	emp1.salary = 47500.00

	emp3 := mycompany{"Manager", "pop", 34500.00}

	fmt.Println("In My company we have")
	fmt.Println("Role: ", emp.role)
	fmt.Println("Name: ", emp.name)
	fmt.Println("salary: ", emp.salary)
	fmt.Println("----------------------------")

	fmt.Println("Role: ", emp1.role)
	fmt.Println("Name: ", emp1.name)
	fmt.Println("salary: ", emp1.salary)
	fmt.Println("----------------------------")

	fmt.Println("Role: ", emp3.role)
	fmt.Println("Name: ", emp3.name)
	fmt.Println("salary: ", emp3.salary)
	fmt.Println("----------------------------")

}

// more?
func QuizMore() {
	// ข้อพิเศษอีก1ข้อ จงเขียนโปรแกรมให้ผลออกมาตามรูปนี้(โดยใช้for)
	// *
	// * *
	// * * *

	for i := 1; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}
