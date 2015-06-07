package main

import "fmt"


func main() {
     var input int
     nextFib := nextFibGenerator()

     fmt.Scanf("%d", &input)

     for i := 0; i < input; i++ {
     	 fmt.Println(nextFib())
     }
}

func nextFibGenerator() func() int {
     cur := 1
     prev := 0

     return func() (ret int) {
     	    ret = prev
	    prev = cur
	    cur = cur + ret

	    return
     }

}