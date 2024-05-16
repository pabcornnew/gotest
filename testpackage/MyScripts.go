package testpackage

import "fmt"

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

func MinSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return arr
}

func MaxSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return arr
}

func Quiz2() {
	//เขียนโปรแกรมหาเลขที่น้อยสุดและมากสุด (เรียกใช้เป็นฟังชั่น) ฝึกการใช้ฟังชั่น (หน้า34-35)
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Printf("Array count : %d\n", len(x))
	fmt.Println("Original array:", x)
	MinSort(x)
	fmt.Println("Min Sorted array:  ", x)
	MaxSort(x)
	fmt.Println("Max Sorted array:  ", x)
	fmt.Println("--------")
}

// 3
func Quiz3() {
	// 1-1000 มีเลข9รวมทั้งหมดกี่ตัว จงเขียนโปรแกรม(คำตอบข้อ3 300)
	// count := 0
	// for i := 1; i <= 1000; i++ {
	// 	if i == "%9%" {

	// 	}
	// }
	// println(count)
}
