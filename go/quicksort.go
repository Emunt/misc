package main


import "fmt"


func main() {

     a := []int{5,1,2,8,5,2,4,5,2,123,3,4,7,6,2,5,1,2,4,5,1}
     QuickSort(a)
     fmt.Println(a)
}

func swap(list []int, i,j int) {
     temp := list[i]
     list[i] = list[j]
     list[j] = temp
}

func QuickSort(list []int) {
     highSwapPos := len(list) - 1
     pivotPos := 0

     for highSwapPos > pivotPos {
     	 if list[pivotPos+1] > list[pivotPos] {
	    swap(list,pivotPos+1, highSwapPos)
	    highSwapPos--
	 } else {
	   swap(list, pivotPos+1, pivotPos)
	   pivotPos++
	 }
     }
     if pivotPos > 0 {
     	QuickSort(list[:pivotPos])
     }

     if pivotPos < len(list) {
     	QuickSort(list[pivotPos+1:])
     }
}
